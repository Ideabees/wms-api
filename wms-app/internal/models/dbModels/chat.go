package dbModels

import (
	"time"

	"gorm.io/gorm"
)

type CreateChats struct {
	gorm.Model
	ChatID               string `gorm:"primaryKey;autoIncrement:false"`
	SenderID             string `gorm:"type:string;not null;uniqueIndex:idx_sender_receiver"`
	ReceiverID           string `gorm:"type:string;not null;uniqueIndex:idx_sender_receiver"`
	ReceiverMobileNumber string
	CreatedAt            time.Time
}

type CreateMessages struct {
	gorm.Model
	MessageID        string `gorm:"primaryKey;autoIncrement:false"`
	ChatID           string `gorm:"column:chat_id;not null"`
	SenderID         string
	MessageType      string // e.g., "text", "image", etc.
	Content          string
	MediaURL         string
	IsReadByReceiver string // "true" or "false"
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
