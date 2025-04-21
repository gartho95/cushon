package database

import "cushonTechTest/models"

func Migrate() error {

	if err := db.AutoMigrate(&models.Type{}, &models.Fund{}, &models.User{}, &models.Account{}); err != nil {
		return err
	}

	return nil
}
