package driver

import (
	"fmt"
	"log"

	"github.com/constellatehq/auth-api/config"
	"github.com/jmoiron/sqlx"
)

func InitDb() (*sqlx.DB, error) {

	sqlxPostgresAddr := buildSqlxPostgresAddr(config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := sqlx.Connect("pgx", sqlxPostgresAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// postgresUrl := BuildPostgresUrl(config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	// Migrate(postgresUrl)

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	// db.MustExec(schema)

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
