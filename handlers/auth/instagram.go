package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/instagram"
)

var (
	instagramOauthConfig *oauth2.Config
	baseInstagramApiUrl = "https://graph.instagram.com"
	instagramClientID string
	instagramClientSecret string
	instagramRedirectUrl = "https://localhost:8000/auth/instagram/callback"
)

func InitInstagramClient() {
	instagramClientID = os.Getenv("INSTAGRAM_CLIENT_ID")
	instagramClientSecret = os.Getenv("INSTAGRAM_CLIENT_SECRET")

	instagramOauthConfig = &oauth2.Config{
		RedirectURL:  instagramRedirectUrl,
		ClientID:     instagramClientID,
		ClientSecret: instagramClientSecret,
		Scopes:       []string{"user_profile", "user_media"},
		Endpoint:     instagram.Endpoint,
	}
}
func InstagramLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := instagramOauthConfig.AuthCodeURL(oauthStateString)

	redirectUrl := RedirectUrl{url}
	json.NewEncoder(w).Encode(redirectUrl)
}

func InstagramCallbackHandler(w http.ResponseWriter, r *http.Request) {
	content, err := getInstagramUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "Content: %s\n", content)
}


func getInstagramUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := instagramOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	fields := "id,username"
	getInstagramUserUrl := baseInstagramApiUrl + "/me?fields=" + fields + "&access_token=" + token.AccessToken

	response, err := http.Get(getInstagramUserUrl)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}