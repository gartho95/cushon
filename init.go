package main

import (
	"cushonTechTest/database"
	"log/slog"

	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	isDatabaseEmpty, err := database.IsDatabaseEmpty(db)
	if err != nil {
		return nil, err
	}

	if isDatabaseEmpty {
		slog.Info("Running migration and seeders as no accounts have been found.")

		err = database.Migrate()
		if err != nil {
			return nil, err
		}
		err = database.Seed(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
