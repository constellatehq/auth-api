package server

import (
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"github.com/constellatehq/auth-api/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
}

func newServer() {
	r := mux.NewRouter()
	routes.initAuthRoutes(r);
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			title := vars["title"]
			page := vars["page"]

			fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":8080", r)
}

func main() {
	init()

	newServer()

}

