package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type FBUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var (
	facebookOauthConfig  *oauth2.Config
	baseFacebookApiUrl   = "https://graph.facebook.com"
	facebookClientID     string
	facebookClientSecret string
	facebookRedirectUrl  = "https://localhost:8000/auth/facebook/callback"
)

func InitFacebookClient() {
	facebookClientID = os.Getenv("FACEBOOK_CLIENT_ID")
	facebookClientSecret = os.Getenv("FACEBOOK_CLIENT_SECRET")

	facebookOauthConfig = &oauth2.Config{
		RedirectURL:  facebookRedirectUrl,
		ClientID:     facebookClientID,
		ClientSecret: facebookClientSecret,
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}
}

func FacebookLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := facebookOauthConfig.AuthCodeURL(oauthStateString)

	redirectUrl := RedirectUrlResponse{url}
	json.NewEncoder(w).Encode(redirectUrl)
}

func FacebookCallbackHandler(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
	}

	code := r.FormValue("code")

	token, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
	}

	getFacebookUserUrl := baseFacebookApiUrl + "/me?access_token=" + url.QueryEscape(token.AccessToken)

	resp, err := http.Get(getFacebookUserUrl)
	if err != nil {
		fmt.Printf("Get: %s\n", err)
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)

	}

	SetAuthorizationCookie(w, token.AccessToken)

	// var data map[string]interface{}
	var fbUser FBUser
	if err := json.Unmarshal(response, &fbUser); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
