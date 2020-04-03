package routes

import (
	"github.com/gorilla/mux"

	"github.com/constellatehq/auth-api/handlers/auth"
	"github.com/constellatehq/auth-api/handlers/user"
	authMiddleware "github.com/constellatehq/auth-api/middleware/auth"
)

func InitRoutes(r *mux.Router) {
	initAuthRoutes(r)
	initUserRoutes(r)
}

func initAuthRoutes(r *mux.Router) *mux.Router {
	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/google", auth.GoogleLoginHandler).Methods("GET")
	authRouter.HandleFunc("/google/callback", auth.GoogleCallbackHandler).Methods("GET")
	authRouter.HandleFunc("/facebook", auth.FacebookLoginHandler).Methods("GET")
	authRouter.HandleFunc("/facebook/callback", auth.FacebookCallbackHandler).Methods("GET")
	authRouter.HandleFunc("/instagram", auth.InstagramLoginHandler).Methods("GET")
	authRouter.HandleFunc("/instagram/callback", auth.InstagramCallbackHandler).Methods("GET")
	authRouter.HandleFunc("/spotify", auth.SpotifyLoginHandler).Methods("GET")
	authRouter.HandleFunc("/spotify/callback", auth.SpotifyCallbackHandler).Methods("GET")

	return authRouter
}

func initUserRoutes(r *mux.Router) *mux.Router {
	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.Use(authMiddleware.GenerateAuthMiddleware())
	userRouter.HandleFunc("/profile", user.UserProfileHandler).Methods("GET")

	return userRouter
}
