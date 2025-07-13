package dbModels

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerId   string `gorm:"primaryKey;autoIncrement:false"`
	FirstName    string
	LastName     string
	MobileNumber string
	UserId       string
	EmailID      string
	City         string
	Pincode      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
