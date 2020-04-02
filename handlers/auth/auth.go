package auth

import (
	"net/http"
	"time"
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
		Name:    "AccessToken",
		Value:   "Bearer " + accessToken,
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
}
