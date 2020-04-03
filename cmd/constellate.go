package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/constellatehq/auth-api/handlers/auth"
	"github.com/constellatehq/auth-api/routes"
	facebookClient "github.com/constellatehq/auth-api/server/clients/facebook_client"
	googleClient "github.com/constellatehq/auth-api/server/clients/google_client"
	"github.com/constellatehq/auth-api/utilities"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	port = "127.0.0.1:8000"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	utilities.InitEnv()

	googleClient.InitClient()
	facebookClient.InitClient()
	auth.InitInstagramClient()
	auth.InitSpotifyClient()
}

func main() {

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},   // All origins
		AllowedMethods: []string{"GET"}, // Allowing only get, just an example
	})

	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))
	routes.InitRoutes(router)
	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		status := Status{200, "Server is healthy"}

		json.NewEncoder(w).Encode(status)
	})

	handler := cors.Handler(router)

	srv := &http.Server{
		Handler: handler,
		Addr:    port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on port:", port)
	log.Fatal(srv.ListenAndServeTLS("https-server.crt", "https-server.key"))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
