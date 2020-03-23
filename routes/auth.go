package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/constellatehq/auth-api/handlers"
)

func initAuthRoutes(r *mux.Router) {
	authRouter.HandleFunc("/auth/google", handlers.googleLogin).Methods("GET")
	authRouter.HandleFunc("/auth/google/redirect", handlers.googleCallback).Methods("GET")
	authRouter := r.PathPrefix("/auth").Subrouter()
}