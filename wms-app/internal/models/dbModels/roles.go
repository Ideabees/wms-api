package dbModels

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	gorm.Model
	EmployeeId   string `gorm:"primaryKey;autoIncrement:false"`
	FirstName    string
	LastName     string
	MobileNumber string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Role         string
	CreatedBy    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastLoggedIn time.Time
}
