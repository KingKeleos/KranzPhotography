package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Setup() (err error) {
	m, err := migrate.New(
		"file:database/migrations/",
		"postgres://localhost:5432/website?sslmode=disable")
	if err != nil {
		return
	}
	err = m.Up()
	return
}
