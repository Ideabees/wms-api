package request

type SendMessageOneToOne struct {
	ChatID               string `json:"chat_id" binding:"required"`
	ReceiverID           string `json:"receiver_id" binding:"required"`
	ReceiverMobileNumber string `json:"receiver_mobile_number" binding:"required"`
	Message              string `json:"message" binding:"required"`
	MessageType          string `json:"message_type" binding:"required"` // e.g., "text", "image", etc.
}

type CreateChat struct {
	ReceiverMobileNumber string `json:"receiver_mobile_number" binding:"required"`
}
