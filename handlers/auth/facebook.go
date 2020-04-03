package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/constellatehq/auth-api/model"
	facebookClient "github.com/constellatehq/auth-api/server/clients/facebook_client"
	fb "github.com/huandu/facebook"
	"golang.org/x/oauth2"
)

func GetFacebookClientID() string {
	return facebookClient.ClientID
}

func FacebookLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := facebookClient.OauthConfig.AuthCodeURL(oauthStateString)

	redirectUrl := RedirectUrlResponse{url}
	json.NewEncoder(w).Encode(redirectUrl)
}

func FacebookCallbackHandler(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		model.CreateErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Invalid OAuth state", nil)
		return
	}

	code := r.FormValue("code")

	token, err := facebookClient.OauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		model.CreateErrorResponse(w, http.StatusBadRequest, "Bad Request", err.Error(), nil)
		return
	}
	fmt.Println("FB Access Token:", url.QueryEscape(token.AccessToken))

	session := facebookClient.GlobalApp.Session(token.AccessToken)

	fields := "id,first_name,last_name,email,gender,age_range,birthday"

	response, err := session.Get("/me", fb.Params{
		"fields": fields,
	})
	if err != nil {
		model.CreateErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "FB Client error: "+err.Error(), nil)
		return
	}

	fmt.Println("Facebook sdk call:", response["id"], response["email"])

	SetAuthorizationCookie(w, token.AccessToken)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
