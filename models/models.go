package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	BaseModel
	FirstName    *string `gorm:"column:first_name"`
	LastName     *string `gorm:"column:last_name"`
	OfficialName *string `gorm:"column:official_name"`
	TypeID       uuid.UUID
	Type         Type `gorm:"foreignKey:TypeID"`
}

type Type struct {
	BaseModel
	Name *string `gorm:"column:name"`
	Code *string `gorm:"column:code"`
}

type Account struct {
	BaseModel
	FundID  uuid.UUID
	UserID  uuid.UUID
	Balance *int
	Fund    Fund `gorm:"foreignKey:FundID"`
	User    User `gorm:"foreignKey:UserID"`
}

type Fund struct {
	BaseModel
	Name *string `gorm:"column:name"`
}
