package main

import (
	"log/slog"

	"github.com/KingKeleos/KranzPhotography/database"
)

func main() {
	err := database.Setup()
	if err != nil {
		slog.Error("migrating the database", "error", err)
	}
	slog.Info("successfully migrated the database")
}
