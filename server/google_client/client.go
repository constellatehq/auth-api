package googleClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/constellatehq/auth-api/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	OAUTH_REDIRECT_URL = "https://localhost:8000/auth/google/callback"
	BASE_API_URL       = "https://www.googleapis.com/oauth2"
)

var (
	OauthConfig  *oauth2.Config
	ClientID     string
	ClientSecret string
)

func InitClient() {
	ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

	OauthConfig = &oauth2.Config{
		RedirectURL:  OAUTH_REDIRECT_URL,
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func Get(url string) (models.Response, error) {
	response, err := http.Get(BASE_API_URL + url)
	if err != nil {
		return nil, fmt.Errorf("Google Get request failed: %s", err.Error())
	}

	defer response.Body.Close()
	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed reading response body: %s", err.Error())
	}

	return makeResponse(buf.Bytes())
}

func makeResponse(response []byte) (models.Response, error) {
	var res models.Response
	err := json.Unmarshal(response, &res)
	if err != nil {
		return nil, fmt.Errorf("Failed converting response to object: %s", err.Error())
	}

	return res, nil

}
