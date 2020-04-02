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
	facebookVerificationUrl = ""
)

func GenerateAuthMiddleware(oauthProvider string) (mw func(http.Handler) http.Handler) {
	mw = func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Use foo1 & foo2
				h.ServeHTTP(w, r)
		})
	}
	return
}