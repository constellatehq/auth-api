package googleClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/constellatehq/auth-api/model"
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
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/user.birthday.read"},
		Endpoint:     google.Endpoint,
	}
}

func Api(url string, requestType string, accessToken string) ([]byte, error) {
	req, err := http.NewRequest(requestType, BASE_API_URL+url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)

	body, err := ResponseToBytes(response)

	if response.StatusCode != 200 {
		return body, fmt.Errorf("%s", response.Status)
	}

	return body, nil
}

func ResponseToBytes(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	buf := &bytes.Buffer{}
	_, err := io.Copy(buf, response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed reading response body: %s", err.Error())
	}
	return buf.Bytes(), nil
}

func Get(url string) ([]byte, error) {
	response, err := http.Get(BASE_API_URL + url)
	if err != nil {
		return nil, fmt.Errorf("Google Get request failed: %s", err.Error())
	}

	body, err := ResponseToBytes(response)

	if response.StatusCode != 200 {
		return body, fmt.Errorf("%s", response.Status)
	}

	return body, err
}

func makeResponse(response []byte) (model.Response, error) {
	var res model.Response
	err := json.Unmarshal(response, &res)
	if err != nil {
		return nil, fmt.Errorf("Failed converting response to object: %s", err.Error())
	}

	return res, nil
}
