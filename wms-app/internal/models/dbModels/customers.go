package dbModels

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	CustomerId   string `gorm:"primaryKey;autoIncrement:false"`
	FirstName    string
	LastName     string
	MobileNumber string `gorm:"unique"`
	CreatedBy    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
