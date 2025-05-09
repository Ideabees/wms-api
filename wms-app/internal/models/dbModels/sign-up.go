package dbModels

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserId       string `gorm:"primaryKey;autoIncrement:false"`
	FirstName    string
	MiddleName   string
	LastName     string
	Email        string `gorm:"unique"`
	PasswordHash string
	MobileNumber string `gorm:"unique"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}