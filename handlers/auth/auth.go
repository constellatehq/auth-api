package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/constellatehq/auth-api/config"
)

type RedirectUrlResponse = struct {
	RedirectUrl string `json:"redirect_url"`
}

var (
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func CreateClientCookie(name string, value string, expiration time.Time) http.Cookie {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Domain:   config.ConstellateDomain,
		Expires:  expiration,
		HttpOnly: false,
		Path:     "/",
	}

	return cookie
}

func SetOauthStateCookie(w http.ResponseWriter, state string) {
	expiration := time.Now().Add(25 * time.Minute)
	cookie := CreateClientCookie("OauthState", state, expiration)

	http.SetCookie(w, &cookie)
}

func SetAuthorizationCookie(w http.ResponseWriter, accessToken string) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := CreateClientCookie("ConstellateAccessToken", fmt.Sprintf("Bearer %s", accessToken), expiration)

	http.SetCookie(w, &cookie)
}
