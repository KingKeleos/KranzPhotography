package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

type Connection struct {
	Username string
	Password string
	Host     string
	SSLMode  string
}

func (c *Connection) Connect() (*sql.DB, error) {
	connection := fmt.Sprintf("postgres://%s:%s@%s:5432?sslmode=%s", c.Username, c.Password, c.Host, c.SSLMode)
	conn, err := sql.Open("postgres", connection)
	if err != nil {
		slog.Error("connecting to the database", "Error", err)
		return nil, err
	}
	return conn, nil
}

func (c *Connection) ReadEnvs() error {
	c.Username = os.Getenv("DB_USERNAME")
	c.Password = os.Getenv("DB_PASSWORD")
	c.Host = os.Getenv("DB_HOST")
	c.SSLMode = os.Getenv("DB_SSL")
	return nil
}
