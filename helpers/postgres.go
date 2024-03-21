package helpers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func BuildPostgresConnUrl() string {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbName)
}

func StartPostgresDb() *sql.DB {
	db, err := sql.Open("postgres", BuildPostgresConnUrl())

	if err != nil {

		panic(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return db
}
