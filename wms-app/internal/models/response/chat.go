package response

import "time"

type GetUserChatsResponse struct {
	Chats []Chat `json:"chats"`
}

type Chat struct {
	ChatID               string    `json:"chat_id"`
	SenderID             string    `json:"sender_id"`
	ReceiverID           string    `json:"receiver_id"`
	ReceiverMobileNumber string    `json:"receiver_mobile_number"`
	CreatedAt            time.Time `json:"created_at"`
}

// get all massages based on chat id
type GetMessagesResponse struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	MessageID   string    `json:"message_id"`
	ChatID     string    `json:"chat_id"`
	SenderID   string    `json:"sender_id"`
	MessageType string    `json:"message_type"`
	Content    string    `json:"content"`
	MediaURL   string    `json:"media_url"`
	IsReadByReceiver string `json:"is_read_by_receiver"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}