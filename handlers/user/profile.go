package user

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Id: "1234", Name: "Felix Zheng"}
	json.NewEncoder(w).Encode(user)
}
