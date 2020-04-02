package facebookClient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	fb "github.com/huandu/facebook"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

const (
	BASE_API_URL       = "https://graph.facebook.com"
	OAUTH_REDIRECT_URL = "https://localhost:8000/auth/facebook/callback"
)

var (
	OauthConfig  *oauth2.Config
	ClientID     string
	ClientSecret string
	GlobalApp    *fb.App
)

func InitFacebookClient() {
	ClientID = os.Getenv("FACEBOOK_CLIENT_ID")
	ClientSecret = os.Getenv("FACEBOOK_CLIENT_SECRET")

	GlobalApp = fb.New(ClientID, ClientSecret)

	OauthConfig = &oauth2.Config{
		RedirectURL:  OAUTH_REDIRECT_URL,
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}
}

func MakeRequest(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Get: %s\n", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)
	}

	return body
}
