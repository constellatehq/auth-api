package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"github.com/constellatehq/auth-api/config"
	"github.com/constellatehq/auth-api/handlers/auth"
	"github.com/constellatehq/auth-api/routes"
	facebookClient "github.com/constellatehq/auth-api/server/clients/facebook"
	googleClient "github.com/constellatehq/auth-api/server/clients/google"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	appPort = "127.0.0.1:8000"
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

func initDb() (*sqlx.DB, error) {

	postgresAddr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := sqlx.Connect("postgres", postgresAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	// db.MustExec(schema)

	return db, nil
}

func main() {

	initDb()

	cors := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedOrigins:   []string{"http://localhost:3000", "https://constellatehq.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		Debug:            true,
	})

	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	routes.InitRoutes(router)
	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		status := Status{200, "Server is healthy"}

		json.NewEncoder(w).Encode(status)
	})

	handler := cors.Handler(router)

	srv := &http.Server{
		Handler: handler,
		Addr:    appPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on port:", appPort)
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
