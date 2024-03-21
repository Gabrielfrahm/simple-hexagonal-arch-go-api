package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"simple-hexagonal-arch-go-api/helpers"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//go:embed sql/*.sql
var migFs embed.FS

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", helpers.BuildPostgresConnUrl())
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	d, err := iofs.New(migFs, "sql")
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	version, _, err := m.Version()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Migration applied: %d\n", version)
}
