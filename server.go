package main

import (
	"log"
	"time"
	"encoding/json"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/gorilla/mux"

	"github.com/constellatehq/auth-api/handlers/auth"
	"github.com/constellatehq/auth-api/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	auth.InitGoogleClient()
	auth.InitFacebookClient()
	auth.InitInstagramClient()
	auth.InitSpotifyClient()
}

type Status struct {
	Code		int			`json:"code"`
	Message	string	`json:"message"`
}

func newServer() {

}

func main() {

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET"}, // Allowing only get, just an example
	})

	newServer()
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))
	routes.InitRoutes(router);
	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		status := Status{200, "Server is healthy"}

    json.NewEncoder(w).Encode(status)
	})

	handler := cors.Handler(router)

	srv := &http.Server{
		Handler:      handler,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServeTLS("https-server.crt", "https-server.key"))
	log.Println("Server started on port 8000")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do stuff here
			log.Println(r.RequestURI)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
	})
}