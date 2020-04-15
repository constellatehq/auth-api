package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/constellatehq/auth-api/config"
	"github.com/constellatehq/auth-api/model"
	"github.com/constellatehq/auth-api/model/errors"
	"github.com/constellatehq/auth-api/model/schema"
	"github.com/constellatehq/auth-api/repository"
	facebookClient "github.com/constellatehq/auth-api/server/clients/facebook"
	fb "github.com/huandu/facebook"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
)

func GetFacebookClientID() string {
	return facebookClient.ClientID
}

func FacebookLoginHandler(w http.ResponseWriter, r *http.Request) {
	oauthState = r.FormValue("state")
	fmt.Printf("State: %s\n", oauthState)
	if oauthState == "" {
		oauthState = oauthStateString
	}
	url := facebookClient.OauthConfig.AuthCodeURL(oauthState)

	redirectUrl := RedirectUrlResponse{url}
	json.NewEncoder(w).Encode(redirectUrl)
}

func FacebookCallbackHandler(env *model.Env, w http.ResponseWriter, r *http.Request) {

	// Check state clientside since we send state as param from client
	state := r.FormValue("state")
	// if state != oauthStateString {
	// 	model.CreateErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Invalid OAuth state", nil)
	// 	return
	// }

	code := r.FormValue("code")

	token, err := facebookClient.OauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		model.CreateErrorResponse(w, http.StatusBadRequest, "Bad Request", err.Error(), nil)
		return
	}
	fmt.Println("FB Access Token:", url.QueryEscape(token.AccessToken))

	_, err = getFacebookUserInfo(env.Db, token.AccessToken)
	if err != nil {
		model.CreateErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "FB Client error: "+err.Error(), nil)
		return
	}

	SetAuthorizationCookie(w, token.AccessToken)
	SetOauthStateCookie(w, state)

	http.Redirect(w, r, config.OauthRedirectUrl, 302)
}

func getFacebookUserInfo(db *sqlx.DB, accessToken string) (*schema.FacebookUserInfoResponse, error) {

	session := facebookClient.GlobalApp.Session(accessToken)

	fields := "id,first_name,last_name,email,gender,age_range,birthday"

	response, err := session.Get("/me", fb.Params{
		"fields": fields,
	})
	if err != nil {
		return nil, err
	}

	// Format FB's birthday date response for time type
	const shortForm = "01/02/2006"
	response["birthday"], _ = time.Parse(shortForm, response["birthday"].(string))

	bytes, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling Facebook user info response: %s\n", err)
	}

	userInfoResponse := schema.FacebookUserInfoResponse{}

	err = json.Unmarshal(bytes, &userInfoResponse)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling Facebook user info response: %s\n", err)
	}

	var user model.User
	user.FacebookId = &userInfoResponse.Id
	user.FirstName = userInfoResponse.FirstName
	user.LastName = userInfoResponse.LastName
	user.Email = userInfoResponse.Email
	user.Birthday = &userInfoResponse.Birthday
	// user.Gender = &userInfoResponse.Gender

	_, err = repository.CreateUserIfNotExists(db, "email", userInfoResponse.Email, user)
	switch err {
	case nil:

	case errors.UserExistsError:

	default:
		return &userInfoResponse, err
	}

	return &userInfoResponse, nil
}
