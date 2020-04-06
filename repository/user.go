package repository

import (
	"github.com/constellatehq/auth-api/model"
	"github.com/jmoiron/sqlx"
)

type UserInterface interface {
	GetUserById(db *sqlx.DB, id string) (*model.User, error)
}

func GetUserById(db *sqlx.DB, id string) (*model.User, error) {
	return
}
