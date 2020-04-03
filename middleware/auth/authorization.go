package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/constellatehq/auth-api/model"
	facebookClient "github.com/constellatehq/auth-api/server/clients/facebook_client"
	googleClient "github.com/constellatehq/auth-api/server/clients/google_client"
	"github.com/gorilla/context"
)

func GenerateAuthMiddleware() (mw func(http.Handler) http.Handler) {
	mw = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken := r.Header.Get("Authorization")
			fmt.Println("\nAccessToken:", accessToken)
			tokenType, accessToken := splitToken(accessToken)

			session := facebookClient.GlobalApp.Session(accessToken)
			googleResponse, googleErr := validateGoogleToken(accessToken)
			if googleErr != nil {
				fmt.Printf("%s\n", googleErr)
			}
			fmt.Printf("%+v\n", googleResponse)

			// Facebook SDK's equivalent of token validation
			facebookResponse, facebookErr := session.Inspect()
			if facebookErr != nil {
				fmt.Printf("Error validating Facebook access token: %s\n", facebookErr)
			}

			if googleErr != nil && facebookErr != nil {
				model.CreateErrorResponse(w, http.StatusUnauthorized, "Unauthorized", "Invalid access token", nil)
				return
			}

			fmt.Printf("%+v\n", facebookResponse)
			// Get user off id and set user/access token
			context.Set(r, "accessToken", tokenType+" "+accessToken)
			next.ServeHTTP(w, r)
		})
	}
	return
}

func validateGoogleToken(accessToken string) (model.Response, error) {
	response, err := googleClient.Api("/v3/userinfo", "GET", accessToken)
	if err != nil {
		return nil, fmt.Errorf("Error validating Google access token: %s\n", err)
	}

	return response, nil
}

func splitToken(accessToken string) (string, string) {
	splitToken := strings.Split(accessToken, " ")

	if len(splitToken) < 2 {
		return "", ""
	}

	return splitToken[0], splitToken[1]
}
