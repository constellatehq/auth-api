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

func SetOauthStateCookie(w http.ResponseWriter, state string) {
	expiration := time.Now().Add(25 * time.Minute)

	cookie := http.Cookie{
		Name:     "OauthState",
		Value:    state,
		Domain:   config.ConstellateDomain,
		Expires:  expiration,
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}

func SetAuthorizationCookie(w http.ResponseWriter, accessToken string) {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	cookie := http.Cookie{
		Name:     "ConstellateAccessToken",
		Value:    fmt.Sprintf("Bearer %s", accessToken),
		Domain:   config.ConstellateDomain,
		Expires:  expiration,
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}
