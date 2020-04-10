package routes

import (
	"github.com/gorilla/mux"

	"github.com/constellatehq/auth-api/handlers/auth"
	"github.com/constellatehq/auth-api/handlers/user"
	authMiddleware "github.com/constellatehq/auth-api/middleware/auth"
	"github.com/constellatehq/auth-api/model"
)

func InitRoutes(r *mux.Router, env model.Env) {
	initAuthRoutes(r, env)
	initUserRoutes(r, env)
}

func initAuthRoutes(r *mux.Router, env model.Env) *mux.Router {
	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/google", auth.GoogleLoginHandler).Methods("GET")
	authRouter.Handle("/google/callback", model.Handler{Env: &env, H: auth.GoogleCallbackHandler}).Methods("GET")
	authRouter.HandleFunc("/facebook", auth.FacebookLoginHandler).Methods("GET")
	authRouter.Handle("/facebook/callback", model.Handler{Env: &env, H: auth.FacebookCallbackHandler}).Methods("GET")
	authRouter.HandleFunc("/instagram", auth.InstagramLoginHandler).Methods("GET")
	authRouter.HandleFunc("/instagram/callback", auth.InstagramCallbackHandler).Methods("GET")
	authRouter.HandleFunc("/spotify", auth.SpotifyLoginHandler).Methods("GET")
	authRouter.HandleFunc("/spotify/callback", auth.SpotifyCallbackHandler).Methods("GET")

	return authRouter
}

func initUserRoutes(r *mux.Router, env model.Env) *mux.Router {
	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.Use(authMiddleware.GenerateAuthMiddleware(env.Db))
	userRouter.Handle("/profile/{id:[0-9]+}", model.Handler{Env: &env, H: user.UserProfileHandler}).Methods("GET")
	userRouter.Handle("/me", model.Handler{Env: &env, H: user.UserMeHandler}).Methods("GET")

	return userRouter
}
