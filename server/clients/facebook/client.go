package facebookClient

import (
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

func InitClient() {
	ClientID = os.Getenv("FACEBOOK_CLIENT_ID")
	ClientSecret = os.Getenv("FACEBOOK_CLIENT_SECRET")

	GlobalApp = fb.New(ClientID, ClientSecret)

	OauthConfig = &oauth2.Config{
		RedirectURL:  OAUTH_REDIRECT_URL,
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Scopes:       []string{"public_profile", "email", "user_birthday", "user_gender", "user_photos"},
		Endpoint:     facebook.Endpoint,
	}
}
