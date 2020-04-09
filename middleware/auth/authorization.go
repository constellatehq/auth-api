package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/constellatehq/auth-api/model"
	"github.com/constellatehq/auth-api/model/schema"
	"github.com/constellatehq/auth-api/repository"
	facebookClient "github.com/constellatehq/auth-api/server/clients/facebook"
	googleClient "github.com/constellatehq/auth-api/server/clients/google"
	"github.com/gorilla/context"
	"github.com/jmoiron/sqlx"
)

func GenerateAuthMiddleware(db *sqlx.DB) (mw func(http.Handler) http.Handler) {
	mw = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken := r.Header.Get("Authorization")
			tokenType, accessToken := splitToken(accessToken)

			googleResponse, googleErr := validateGoogleToken(accessToken)

			// Facebook SDK's equivalent of token validation
			facebookResponse, facebookErr := validateFacebookToken(accessToken)
			var user model.User
			var err error

			if googleErr != nil && facebookErr != nil {
				model.CreateErrorResponse(w, http.StatusUnauthorized, "Unauthorized", "Invalid access token", nil)
				return
			} else if googleErr != nil && facebookErr == nil {
				user, err = repository.GetUserByField(db, "facebook_id", facebookResponse.UserId)
				if err != nil {
					model.CreateErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), nil)
					return
				}
			} else if facebookErr != nil && googleErr == nil {
				user, err = repository.GetUserByField(db, "google_id", googleResponse.Sub)
				if err != nil {
					model.CreateErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), nil)
					return
				}
			}

			// Get user off id and set user/access token
			context.Set(r, "accessToken", tokenType+" "+accessToken)
			context.Set(r, "user", user)
			next.ServeHTTP(w, r)
		})
	}
	return
}

func validateGoogleToken(accessToken string) (*schema.GoogleTokenValidationResponse, error) {
	response, err := googleClient.Api("/v3/userinfo", "GET", accessToken)
	if err != nil {
		return nil, fmt.Errorf("Error validating Google access token: %s\n", err)
	}

	tokenValidationResponse := schema.GoogleTokenValidationResponse{}

	err = json.Unmarshal(response, &tokenValidationResponse)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling Google token response: %s\n", err)
	}

	return &tokenValidationResponse, nil
}

func validateFacebookToken(accessToken string) (*schema.FacebookTokenValidationResponse, error) {

	session := facebookClient.GlobalApp.Session(accessToken)
	response, err := session.Inspect()
	if err != nil {
		return nil, fmt.Errorf("Error validating Facebook access token: %s\n", err)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling Facebook token response: %s\n", err)
	}

	tokenValidationResponse := schema.FacebookTokenValidationResponse{}

	err = json.Unmarshal(bytes, &tokenValidationResponse)
	if err != nil {
		return nil, fmt.Errorf("Error unmarhsaling Facebook token response: %s\n", err)
	}

	return &tokenValidationResponse, nil
}

func splitToken(accessToken string) (string, string) {
	splitToken := strings.Split(accessToken, " ")

	if len(splitToken) < 2 {
		return "", ""
	}

	return splitToken[0], splitToken[1]
}
