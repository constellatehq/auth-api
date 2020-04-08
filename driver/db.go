package driver

import (
	"fmt"
	"log"

	"github.com/constellatehq/auth-api/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitDb() (*sqlx.DB, error) {

	sqlxPostgresAddr := buildSqlxPostgresAddr(config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := sqlx.Connect("pgx", sqlxPostgresAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// facebookId := "1234"
	// gender := model.Gender(1)
	// user := model.User{GoogleId: &facebookId, FirstName: "Joey", LastName: "Blumpkin", Gender: gender.String()}
	// repository.CreateUser(db, user)

	return db, nil
}

func buildSqlxPostgresAddr(user string, password string, host string, port int, database string) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
}

func buildPostgresUrl(user string, password string, host string, port int, database string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, database)
}
