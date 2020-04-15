package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/constellatehq/auth-api/config"
	"github.com/constellatehq/auth-api/model"
	"github.com/constellatehq/auth-api/model/errors"
	"github.com/constellatehq/auth-api/model/schema"
	"github.com/constellatehq/auth-api/repository"
	googleClient "github.com/constellatehq/auth-api/server/clients/google"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
)

var (
	oauthState = ""
)

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	oauthState = r.FormValue("state")
	fmt.Printf("State: %s\n", oauthState)
	if oauthState == "" {
		oauthState = oauthStateString
	}

	url := googleClient.OauthConfig.AuthCodeURL(oauthState)

	redirectUrl := RedirectUrlResponse{url}

	json.NewEncoder(w).Encode(redirectUrl)
}

func GoogleCallbackHandler(env *model.Env, w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	code := r.FormValue("code")

	// Check state clientside since we send state as param from client
	// if state != oauthState {
	// 	model.CreateErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Invalid OAuth state", nil)
	// 	return
	// }

	token, err := googleClient.OauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("+%v\n", token)
		model.CreateErrorResponse(w, http.StatusBadRequest, "Bad Request", err.Error(), nil)
		return
	}

	fmt.Printf("Google Access Token: %s\n", token.AccessToken)

	response, err := getGoogleUserInfo(env.Db, token.AccessToken)
	switch err {
	case nil:

	case errors.UserExistsError:

	default:
		model.CreateErrorResponse(w, http.StatusUnauthorized, "Unauthorized", err.Error(), nil)
		return
	}

	SetAuthorizationCookie(w, token.AccessToken)
	SetOauthStateCookie(w, state)

	fmt.Printf("+%v\n", response)

	http.Redirect(w, r, config.OauthRedirectUrl, 302)
}

func getGoogleUserInfo(db *sqlx.DB, accessToken string) (*schema.GoogleUserInfoResponse, error) {

	response, err := googleClient.Get("/v2/userinfo?access_token=" + accessToken)

	if err != nil {
		return nil, fmt.Errorf("Failed getting user info: %s", err.Error())
	}

	userInfoResponse := schema.GoogleUserInfoResponse{}

	err = json.Unmarshal(response, &userInfoResponse)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling Google token response: %s\n", err)
	}

	var user model.User
	user.GoogleId = &userInfoResponse.Id
	user.FirstName = userInfoResponse.GivenName
	user.LastName = userInfoResponse.FamilyName
	user.Email = userInfoResponse.Email

	fmt.Printf("User: %v\n", user)

	_, err = repository.CreateUserIfNotExists(db, "email", userInfoResponse.Email, user)

	if err != nil {
		return &userInfoResponse, err
	}

	return &userInfoResponse, nil
}
