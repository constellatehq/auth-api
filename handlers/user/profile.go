package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/constellatehq/auth-api/model"
	"github.com/constellatehq/auth-api/repository"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func UserProfileHandler(env *model.Env, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Printf("User id from param: %s\n", id)
	user, err := repository.GetUserById(env.Db, "7c74ba9a-ebb1-47f0-b8e2-d33136b557df")
	if err != nil {
		model.CreateErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", err.Error(), nil)
		return
	}

	fmt.Printf("User: +%v\n", user)
	json.NewEncoder(w).Encode(user)
}

func UserMeHandler(env *model.Env, w http.ResponseWriter, r *http.Request) {
	user, ok := context.GetOk(r, "user")
	if !ok {
		model.CreateErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "Could not get user for provided authorization", nil)
		return
	}

	json.NewEncoder(w).Encode(user)
}
