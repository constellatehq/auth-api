package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/constellatehq/auth-api/server/facebookClient"
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
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
	}

	code := r.FormValue("code")

	token, err := facebookClient.OauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
	}
	fmt.Println("FB Access Token:", url.QueryEscape(token.AccessToken))

	session := facebookClient.GlobalApp.Session(token.AccessToken)

	fields := "id,email,first_name,last_name"

	response, err := session.Get("/me", fb.Params{
		"fields": fields,
	})
	if err != nil {
		fmt.Printf("FB Client error: %s\n", err)
	}

	fmt.Println("Facebook sdk call:", response["id"], response["email"])

	SetAuthorizationCookie(w, token.AccessToken)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
