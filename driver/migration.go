package driver

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	DB_DRIVER      = "postgres"
	MIGRATION_PATH = "file://driver/migrations"
)

func Migrate(postgresUrl string) {
	db, err := sql.Open(DB_DRIVER, postgresUrl)
	if err != nil {
		fmt.Printf(err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Printf(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		MIGRATION_PATH,
		DB_DRIVER, driver)
	if err != nil {
		fmt.Printf(err.Error())
	}

	m.Up()

	db.Close()
}
