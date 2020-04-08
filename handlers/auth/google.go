package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/constellatehq/auth-api/model"
	googleClient "github.com/constellatehq/auth-api/server/clients/google"
	"golang.org/x/oauth2"
)

var (
	oauthState = ""
)

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	oauthState = r.FormValue("state")
	fmt.Printf("State: %s\n", oauthState)
	url := googleClient.OauthConfig.AuthCodeURL(oauthState)

	redirectUrl := RedirectUrlResponse{url}
	json.NewEncoder(w).Encode(redirectUrl)
}

func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	code := r.FormValue("code")

	if state != oauthState {
		model.CreateErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Invalid OAuth state", nil)
		return
	}

	token, err := googleClient.OauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		model.CreateErrorResponse(w, http.StatusBadRequest, "Bad Request", err.Error(), nil)
		return
	}

	response, err := getGoogleUserInfo(token.AccessToken)
	if err != nil {
		model.CreateErrorResponse(w, http.StatusUnauthorized, "Unauthorized", err.Error(), nil)
		return
	}

	SetAuthorizationCookie(w, token.AccessToken)
	SetOauthStateCookie(w, state)
	// w.Header().Set("Content-Type", "application/json")

	fmt.Printf("%s", response)
	// json.NewEncoder(w).Encode(response)
	http.Redirect(w, r, "http://localhost:3000/oauth/callback", 302)
}

func getGoogleUserInfo(accessToken string) (model.Response, error) {

	response, err := googleClient.Get("/v2/userinfo?access_token=" + accessToken)

	if err != nil {
		return nil, fmt.Errorf("Failed getting user info: %s", err.Error())
	}

	return response, nil
}
