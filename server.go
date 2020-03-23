package main

import (
	"log"
	"time"
	"encoding/json"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	// "github.com/gorilla/handlers"
	"github.com/constellatehq/auth-api/routes"
	"github.com/constellatehq/auth-api/handlers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	handlers.InitGoogleClient()
}

type Status struct {
	Code		int			`json:"code"`
	Message	string	`json:"message"`
}

func newServer() {

}

func main() {

	newServer()
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	routes.InitRoutes(r);
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		status := Status{200, "Server is healthy"}

    json.NewEncoder(w).Encode(status)
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do stuff here
			log.Println(r.RequestURI)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
	})
}