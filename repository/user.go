package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/constellatehq/auth-api/model"
	"github.com/constellatehq/auth-api/model/errors"
	"github.com/constellatehq/auth-api/server/utilities"
	gsb "github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type UserInterface interface {
	CreateUser(db *sqlx.DB, user model.User) (string, error)
	GetUserById(db *sqlx.DB, id string) (*model.User, error)
}

var userStruct = gsb.NewStruct(new(model.User))

func CreateUser(db *sqlx.DB, user model.User) (string, error) {
	numId, err := utilities.Sonyflake.NextID()
	if err != nil {
		fmt.Printf("Id generation failed\n")
	}
	id := strconv.FormatUint(numId, 10)
	// id := uuid.NewV4().String()

	ib := gsb.PostgreSQL.NewInsertBuilder()

	ib.InsertInto("users")
	ib.Cols(
		"id",
		"facebook_id",
		"google_id",
		"instagram_id",
		"spotify_id",
		"first_name",
		"last_name",
		"email",
		"birthday",
		"gender",
		"onboarded",
		"permission_level",
		"email_verified",
	)
	ib.Values(
		id,
		user.FacebookId,
		user.GoogleId,
		user.InstagramId,
		user.SpotifyId,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Birthday,
		user.Gender,
		user.Onboarded,
		user.PermissionLevel,
		user.EmailVerified,
	)

	// Execute the query.
	sql, args := ib.Build()
	fmt.Printf("Executing %s\n%s\n", sql, args)

	err = Transact(db, func(tx *sqlx.Tx) error {
		if _, err := tx.Exec(sql, args...); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return id, nil
}

func GetUserById(db *sqlx.DB, userId string) (model.User, error) {
	sb := gsb.PostgreSQL.NewSelectBuilder()

	sb.Select(
		"id",
		"facebook_id",
		"google_id",
		"instagram_id",
		"spotify_id",
		"first_name",
		"last_name",
		"email",
		"birthday",
		"gender",
		"onboarded",
		"permission_level",
		"email_verified",
		"created_at",
		"updated_at",
	)
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

func GetUserByField(db *sqlx.DB, fieldName string, fieldValue string) (model.User, error) {
	sb := gsb.PostgreSQL.NewSelectBuilder()

	sb.Select(
		"id",
		"facebook_id",
		"google_id",
		"instagram_id",
		"spotify_id",
		"first_name",
		"last_name",
		"email",
		"birthday",
		"gender",
		"onboarded",
		"permission_level",
		"email_verified",
		"created_at",
		"updated_at",
	)
	sb.From("users")
	sb.Where(sb.Equal(fieldName, fieldValue))

	sql, args := sb.Build()
	fmt.Printf("%s\n%s\n", sql, args)

	user := model.User{}
	err := db.Get(&user, sql, args...)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUserIfNotExists(db *sqlx.DB, fieldName string, fieldValue string, user model.User) (string, error) {
	existingUser, err := GetUserByField(db, fieldName, fieldValue)
	switch err {
	case nil:
		return existingUser.Id, errors.UserExistsError
	case sql.ErrNoRows:
		return CreateUser(db, user)
	default:
		return "", err
	}

}
