package repository

import (
	"github.com/constellatehq/auth-api/model"
	sb "github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type UserInterface interface {
	GetUserById(db *sqlx.DB, id string) (*model.User, error)
}

var userStruct = sb.NewStruct(new(model.User))

func GetUserById(db *sqlx.DB, id string) (*model.User, error) {
	sb := userStruct.InsertInto("users")

	// Execute the query.
	sql, args := sb.Build()
	return
}
