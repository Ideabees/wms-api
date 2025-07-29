package dbModels

import (
	"time"

	"gorm.io/gorm"
)

type Chats struct {
	gorm.Model
	ChatID   string `gorm:"primaryKey;autoIncrement:false"`
	UserID   string
	ToSenderMobileNumber string
	CustomerID string
	ToReceiverMobileNumber string
	Message  string
	Timestamp time.Time
}