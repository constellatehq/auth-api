package auth

import (
	"fmt"
	"log"
	"time"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var (
	spotifyOauthConfig *oauth2.Config
	baseSpotifyApiUrl = "https://api.spotify.com/v1"
	spotifyClientID string
	spotifyClientSecret string
	spotifyRedirectUrl = "https://localhost:8000/auth/spotify/callback"
)

func InitSpotifyClient() {
	spotifyClientID = os.Getenv("SPOTIFY_CLIENT_ID")
	spotifyClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")

	spotifyOauthConfig = &oauth2.Config{
		RedirectURL:  spotifyRedirectUrl,
		ClientID:     spotifyClientID,
		ClientSecret: spotifyClientSecret,
		Scopes:       []string{"user-read-private", "user-read-email"},
		Endpoint:     spotify.Endpoint,
	}
}
func SpotifyLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := spotifyOauthConfig.AuthCodeURL(oauthStateString)

	redirectUrl := RedirectUrlResponse{url}
	json.NewEncoder(w).Encode(redirectUrl)
}

func SpotifyCallbackHandler(w http.ResponseWriter, r *http.Request) {
	content, err := getSpotifyUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "Content: %s\n", content)
}


func getSpotifyUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := spotifyOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	getSpotifyUserUrl := baseSpotifyApiUrl + "/me"

	contents, err := makeAuthRequest(getSpotifyUserUrl, token.AccessToken)

	return contents, nil
}

func makeAuthRequest(url string, accessToken string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Authorization", "Bearer " + accessToken)

	client := &http.Client{Timeout: time.Second * 10}

	response, err := client.Do(req)

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