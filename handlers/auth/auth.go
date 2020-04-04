package auth

import (
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

func SetAuthorizationCookie(w http.ResponseWriter, accessToken string) {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	cookie := http.Cookie{
		Name:    "ConstellateAccessToken",
		Value:   "Bearer " + accessToken,
		Domain:  config.ConstellateDomain,
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
}
