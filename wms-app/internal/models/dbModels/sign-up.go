package dbModels

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string
	MiddleName   string
	LastName     string
	Email        string `gorm:"unique"`
	PasswordHash string
	MobileNumber string `gorm:"unique"`
}