package auth

import (
	"fmt"
	"os"
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
  "encoding/json"

  "golang.org/x/oauth2"
  "golang.org/x/oauth2/facebook"
)

var (
  facebookOauthConfig *oauth2.Config
  baseFacebookApiUrl = "https://graph.facebook.com"
  facebookClientID string
  facebookClientSecret string
  facebookRedirectUrl = "https://localhost:8000/auth/facebook/callback"
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

	redirectUrl := RedirectUrl{url}
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

  log.Printf("parseResponseBody: %s\n", string(response))
}