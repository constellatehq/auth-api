package routes

import (
	"github.com/gorilla/mux"
	"github.com/constellatehq/auth-api/auth"
)

func InitRoutes(r *mux.Router) {
	initAuthRoutes(r)
}

func initAuthRoutes(r *mux.Router) {
	// authRouter := r.PathPrefix("/auth").Subrouter()
	r.HandleFunc("/auth/google", auth.GoogleLoginHandler).Methods("GET")
	r.HandleFunc("/auth/google/callback", auth.GoogleCallbackHandler).Methods("GET")
	r.HandleFunc("/auth/facebook", auth.GoogleLoginHandler).Methods("GET")
	r.HandlerFunc("/auth/facebook/callback", auth.FacebookCallbackHandler).Methods("GET")
}