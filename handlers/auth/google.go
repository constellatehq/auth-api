package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/constellatehq/auth-api/model"
	googleClient "github.com/constellatehq/auth-api/server/google_client"
	"golang.org/x/oauth2"
)

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := googleClient.OauthConfig.AuthCodeURL(oauthStateString)

	redirectUrl := RedirectUrlResponse{url}
	json.NewEncoder(w).Encode(redirectUrl)
}

func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	response, err := getGoogleUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		model.CreateErrorResponse(w, http.StatusUnauthorized, "Unauthorized", err.Error(), nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getGoogleUserInfo(state string, code string) (model.Response, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("Invalid oauth state")
	}

	token, err := googleClient.OauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("Code exchange failed: %s", err.Error())
	}

	response, err := googleClient.Get("/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		return nil, fmt.Errorf("Failed getting user info: %s", err.Error())
	}

	return response, nil
}
