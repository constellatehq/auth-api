package repository

import (
	"fmt"

	"github.com/constellatehq/auth-api/model"
	gsb "github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type UserInterface interface {
	CreateUser(db *sqlx.DB, user model.User) (string, error)
	GetUserById(db *sqlx.DB, id string) (*model.User, error)
}

var userStruct = gsb.NewStruct(new(model.User))

// , (*model.User, error)
func CreateUser(db *sqlx.DB, user model.User) {

	id := uuid.NewV4().String()

	ib := gsb.PostgreSQL.NewInsertBuilder()

	ib.InsertInto("users")
	ib.Cols("id", "facebook_id", "google_id", "instagram_id", "spotify_id", "first_name", "last_name", "email", "birthday", "gender", "onboarded", "permission_level")
	ib.Values(id, user.FacebookId, user.GoogleId, user.InstagramId, user.SpotifyId, user.FirstName, user.LastName, user.Email, user.Birthday, user.Gender, user.Onboarded, user.PermissionLevel)

	// Execute the query.
	sql, args := ib.Build()
	fmt.Printf("%s\n%s\n", sql, args)

	tx := db.MustBegin()
	tx.MustExec(sql, args...)
	tx.Commit()
	// return
}

func GetUserById(db *sqlx.DB, userId string) (model.User, error) {
	sb := gsb.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "facebook_id", "google_id", "instagram_id", "spotify_id", "first_name", "last_name", "email", "birthday", "gender", "onboarded", "permission_level", "created_at", "updated_at")
	sb.From("users")
	sb.Where(sb.Equal("id", userId))

	sql, args := sb.Build()
	fmt.Printf("%s\n%s\n", sql, args)

	user := model.User{}
	err := db.Get(&user, sql, args...)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return user, fmt.Errorf("Error retrieving user with id: %s. %s", userId, err)
	}

	return user, nil
}
