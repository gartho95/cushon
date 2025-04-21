package database

import (
	"cushonTechTest/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	now := time.Now()

	retailTypeID := uuid.New()
	employerTypeID := uuid.New()

	types := []models.Type{
		{
			BaseModel: models.BaseModel{ID: retailTypeID, CreatedAt: now, UpdatedAt: now},
			Name:      strPtr("retail"),
			Code:      strPtr("RTL"),
		},
		{
			BaseModel: models.BaseModel{ID: employerTypeID, CreatedAt: now, UpdatedAt: now},
			Name:      strPtr("employer"),
			Code:      strPtr("EMP"),
		},
	}
	if err := db.Create(&types).Error; err != nil {
		return err
	}

	userIDs := []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}
	users := []models.User{
		{
			BaseModel:    models.BaseModel{ID: userIDs[0], CreatedAt: now, UpdatedAt: now},
			FirstName:    strPtr("John"),
			LastName:     strPtr("Doe"),
			OfficialName: strPtr("John D."),
			TypeID:       retailTypeID,
		},
		{
			BaseModel:    models.BaseModel{ID: userIDs[1], CreatedAt: now, UpdatedAt: now},
			FirstName:    strPtr("Gareth"),
			LastName:     strPtr("Thomas"),
			OfficialName: strPtr("Gareth T."),
			TypeID:       retailTypeID,
		},
		{
			BaseModel:    models.BaseModel{ID: userIDs[2], CreatedAt: now, UpdatedAt: now},
			FirstName:    strPtr("Grayson"),
			LastName:     strPtr("Sponge"),
			OfficialName: strPtr("Grayson S."),
			TypeID:       retailTypeID,
		},
	}
	if err := db.Create(&users).Error; err != nil {
		return err
	}

	fundIDs := []uuid.UUID{uuid.New(), uuid.New(), uuid.New(), uuid.New()}
	funds := []models.Fund{
		{
			BaseModel: models.BaseModel{ID: fundIDs[0], CreatedAt: now, UpdatedAt: now},
			Name:      strPtr("Cushon Equities Fund."),
		},
		{
			BaseModel: models.BaseModel{ID: fundIDs[1], CreatedAt: now, UpdatedAt: now},
			Name:      strPtr("Small Fund"),
		},
		{
			BaseModel: models.BaseModel{ID: fundIDs[2], CreatedAt: now, UpdatedAt: now},
			Name:      strPtr("Large Fund"),
		},
		{
			BaseModel: models.BaseModel{ID: fundIDs[3], CreatedAt: now, UpdatedAt: now},
			Name:      strPtr("Medium Fund"),
		},
	}
	if err := db.Create(&funds).Error; err != nil {
		return err
	}

	accountIDs := []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}
	balance := 1000
	accounts := []models.Account{
		{
			BaseModel: models.BaseModel{ID: accountIDs[0], CreatedAt: now, UpdatedAt: now},
			FundID:    fundIDs[0],
			UserID:    userIDs[0],
			Balance:   &balance,
		},
		{
			BaseModel: models.BaseModel{ID: accountIDs[1], CreatedAt: now, UpdatedAt: now},
			FundID:    fundIDs[1],
			UserID:    userIDs[1],
			Balance:   &balance,
		},
		{
			BaseModel: models.BaseModel{ID: accountIDs[2], CreatedAt: now, UpdatedAt: now},
			FundID:    fundIDs[2],
			UserID:    userIDs[2],
			Balance:   &balance,
		},
	}
	if err := db.Create(&accounts).Error; err != nil {
		return err
	}

	return nil
}

func strPtr(s string) *string {
	return &s
}
