package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Setup will add the migration files to the database
func Setup() (err error) {
	m, err := migrate.New(
		"file:database/migrations/",
		"postgres://localhost:5432/website?sslmode=disable")
	if err != nil {
		return
	}
	err = m.Up()
	m.Close()
	return
}
