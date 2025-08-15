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
