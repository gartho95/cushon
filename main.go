package main

import (
	"cushonTechTest/api"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error(".env file not loaded", "error", err)
		os.Exit(1)

	}
	db, err := InitDB()
	if err != nil {
		slog.Error("error initializing DB", "error", err)
		os.Exit(2)
	}
	server := api.API{DB: db}
	if err := server.RunWebServer(); err != nil {
		slog.Error("error running server", "error", err)
		os.Exit(3)
	}
}
