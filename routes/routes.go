package routes

import (
	"github.com/gorilla/mux"
	"github.com/constellatehq/auth-api/handlers"
)

func InitRoutes(r *mux.Router) {
	initAuthRoutes(r)
}

func initAuthRoutes(r *mux.Router) {
	// authRouter := r.PathPrefix("/auth").Subrouter()
	r.HandleFunc("/auth/google", handlers.GoogleLoginHandler).Methods("GET")
	r.HandleFunc("/auth/google/callback", handlers.GoogleCallbackHandler).Methods("GET")
}