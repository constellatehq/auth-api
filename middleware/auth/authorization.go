package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/constellatehq/auth-api/server/facebookClient"
	"github.com/gorilla/context"
)

var (
	facebookVerificationUrl = ""
)

type FacebookAppTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware1")
		next.ServeHTTP(w, r)
	})
}

func GenerateAuthMiddleware() (mw func(http.Handler) http.Handler) {
	mw = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken := r.Header.Get("Authorization")
			fmt.Println("\nAccessToken:", accessToken)
			tokenType, accessToken := splitToken(accessToken)

			session := facebookClient.GlobalApp.Session(accessToken)

			// Validate access token
			result, err := session.Inspect()
			if err != nil {
				fmt.Printf("Error validating Facebook access token: %s\n", err)
			}

			fmt.Printf("%+v\n", result)
			// Get user off id and set user/access token
			context.Set(r, "accessToken", tokenType+" "+accessToken)
			next.ServeHTTP(w, r)
		})
	}
	return
}

func splitToken(accessToken string) (string, string) {
	splitToken := strings.Split(accessToken, " ")

	if len(splitToken) < 2 {
		return "", ""
	}

	return splitToken[0], splitToken[1]
}
