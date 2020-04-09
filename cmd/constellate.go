package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/constellatehq/auth-api/config"
	"github.com/constellatehq/auth-api/driver"
	"github.com/constellatehq/auth-api/handlers/auth"
	"github.com/constellatehq/auth-api/model"
	"github.com/constellatehq/auth-api/routes"
	facebookClient "github.com/constellatehq/auth-api/server/clients/facebook"
	googleClient "github.com/constellatehq/auth-api/server/clients/google"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	APP_PORT = "localhost:8000"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	config.InitEnv()

	googleClient.InitClient()
	facebookClient.InitClient()
	auth.InitInstagramClient()
	auth.InitSpotifyClient()
}

func main() {

	db, err := driver.InitDb()
	if err != nil {
		fmt.Printf("DB failed to connect: %s", err)
		return
	}

	cors := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedOrigins:   []string{"http://127.0.0.1:3000", "http://localhost:3000", "https://constellatehq.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		Debug:            true, // Uncomment if we want to debug request headers
	})

	env := model.Env{Db: db}

	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	routes.InitRoutes(router, env)
	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		status := Status{200, "Server is healthy"}

		json.NewEncoder(w).Encode(status)
	})

	handler := cors.Handler(router)

	srv := &http.Server{
		Handler: handler,
		Addr:    APP_PORT,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on port:", APP_PORT)
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
