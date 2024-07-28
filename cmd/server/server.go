package main

import (
	"log/slog"

	"github.com/KingKeleos/KranzPhotography/database"
)

func main() {
	slog.Info("starting backend")
	err := database.Setup()
	if err != nil {
		slog.Error("migrating the database", "error", err)
	}
	slog.Info("successfully migrated the database")
	c := database.Connection{}
	c.ReadEnvs()
	err = c.Connect()
	if err != nil {
		slog.Error("creating a connection to the database")
	}
}
