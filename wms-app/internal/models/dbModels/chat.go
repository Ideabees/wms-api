package dbModels

import (
	"time"
	"gorm.io/gorm"
)

type CreateChats struct {
	gorm.Model
	ChatID               string `gorm:"primaryKey;autoIncrement:false"`
	SenderID             string
	ReceiverID           string
	ReceiverMobileNumber string
	Created_At           time.Time
}

type CreateMessages struct {
	gorm.Model
	MessageID        string `gorm:"primaryKey;autoIncrement:false"`
	ChatID           string
	SenderID         string
	MessageType      string // e.g., "text", "image", etc.
	Content          string
	MediaURL         string
	IsReadByReceiver string // "true" or "false"
	Created_At       time.Time
	Updated_At       time.Time
	Deleted_At       time.Time
}
