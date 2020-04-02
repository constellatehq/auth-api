package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/constellatehq/auth-api/server/facebookClient"
	fb "github.com/huandu/facebook"
	"golang.org/x/oauth2"
)

type FBUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

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

	res, err := session.Get("/me", fb.Params{
		"fields": "id,first_name,last_name",
	})

	if err != nil {
		fmt.Printf("FB Client error: %s\n", err)
	}

	fmt.Println("### Facebook sdk call:", res["id"], res["first_name"])

	getFacebookUserUrl := facebookClient.BASE_API_URL + "/me?access_token=" + url.QueryEscape(token.AccessToken)

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
