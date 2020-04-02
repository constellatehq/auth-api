package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
			tokenType, accessToken := getTokenProvider(accessToken)
			facebookToken := FacebookAppTokenResponse{AccessToken: accessToken, TokenType: tokenType}

			validateFBToken(facebookToken.AccessToken)

			context.Set(r, "accessToken", accessToken)
			next.ServeHTTP(w, r)
		})
	}
	return
}

func validateFBToken(userAccessToken string) {
	facebookClientID := os.Getenv("FACEBOOK_CLIENT_ID")
	facebookClientSecret := os.Getenv("FACEBOOK_CLIENT_SECRET")
	appTokenUrl := "https://graph.facebook.com/oauth/access_token?client_id=" + facebookClientID + "&client_secret=" + facebookClientSecret + "&grant_type=client_credentials"

	appTokenResponse := facebookClient.MakeRequest(appTokenUrl)
	fmt.Println("\nApp token res:", string(appTokenResponse))
	var facebookAppTokenResponse FacebookAppTokenResponse
	if err := json.Unmarshal(appTokenResponse, &facebookAppTokenResponse); err != nil {
		panic(err)
	}

	debugTokenUrl := "https://graph.facebook.com/debug_token?input_token=" + userAccessToken + "&access_token=" + facebookAppTokenResponse.AccessToken
	fmt.Println("\nDebug token url:", debugTokenUrl)
	debugTokenResponse := facebookClient.MakeRequest(debugTokenUrl)

	fmt.Printf(string(debugTokenResponse))
}

func getTokenProvider(accessToken string) (string, string) {
	splitToken := strings.Split(accessToken, " ")

	if len(splitToken) < 2 {
		return "", ""
	}

	return splitToken[0], splitToken[1]
}
