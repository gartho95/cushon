package database

import (
	"cushonTechTest/models"

	"gorm.io/gorm"
)

func IsDatabaseEmpty(db *gorm.DB) (bool, error) {
	if !db.Migrator().HasTable(&models.Account{}) {
		return true, nil
	}

	var count int64
	if err := db.Model(&models.Account{}).Count(&count).Error; err != nil {
		return true, err
	}
	return count == 0, nil
}
