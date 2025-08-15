package services

import (
	"fmt"
	"wms-app/config"
	"wms-app/internal/models/dbModels"
	"wms-app/internal/models/response"
	thirdparty "wms-app/internal/third-party"
)

func CreateChat(createChat *dbModels.CreateChats) (string, error) {
	// Insert customer into the database

	// get customer ID by customer mobile number
	var customerID string
	if createChat.ReceiverMobileNumber != "" {
		customerID = getCustomerID(createChat.ReceiverMobileNumber)
	}

	createChat.ReceiverID = customerID

	result := config.DB.Create(createChat)
	if result.Error != nil {
		fmt.Println("Error in create user repository")
		return "Failed", result.Error
	}
	return createChat.ChatID, nil
}

func GetChat(chatID string) (map[string]interface{}, error) {
	// TODO: Implement chat retrieval logic
	return map[string]interface{}{"chat_id": chatID}, nil
}
func GetUserChats(userID string) (response.GetUserChatsResponse, error) {
	// TODO: Implement user chats retrieval logic
	// fetch from db
	var chats []dbModels.CreateChats
	result := config.DB.Where("sender_id = ?", userID).Find(&chats)
	if result.Error != nil {
		return response.GetUserChatsResponse{}, result.Error
	}

	var responseChats []response.Chat
	for _, chat := range chats {
		responseChats = append(responseChats, response.Chat{
			ChatID:               chat.ChatID,
			SenderID:             chat.SenderID,
			ReceiverID:           chat.ReceiverID,
			ReceiverMobileNumber: chat.ReceiverMobileNumber,
			CreatedAt:            chat.CreatedAt,
		})
	}

	return response.GetUserChatsResponse{Chats: responseChats}, nil
}
func SendMessage(msg *dbModels.CreateMessages) (string, error) {
	result := config.DB.Create(msg)
	if result.Error != nil {
		return "", result.Error
	}
	// send msg to third-party service

	return msg.MessageID, nil
}


func GetMessages(chatID string) (response.GetMessagesResponse, error) {

	var messages []dbModels.CreateMessages

	result := config.DB.Where("chat_id = ?", chatID).Find(&messages)
	if result.Error != nil {
		fmt.Println("Error retrieving messages for chat ID:", chatID, "Error:", result.Error)
		return response.GetMessagesResponse{}, result.Error
	}
	fmt.Println("Messages retrieved successfully for chat ID:", chatID)

	var responseMessages []response.Message

	fmt.Println("Processing messages for chat ID:", chatID)
	if len(messages) == 0 {
		fmt.Println("No messages found for chat ID:", chatID)
	} else {
		fmt.Println("Messages found for chat ID:", chatID)
	}
	// Process each message and map to response format

	for _, msg := range messages {
		fmt.Println("Processing message:", msg.MessageID)
		responseMessages = append(responseMessages, response.Message{
			MessageID:        msg.MessageID,
			ChatID:           msg.ChatID,
			SenderID:         msg.SenderID,
			MessageType:      msg.MessageType,
			Content:          msg.Content,
			MediaURL:         msg.MediaURL,
			IsReadByReceiver: msg.IsReadByReceiver,
			CreatedAt:        msg.CreatedAt,
			UpdatedAt:        msg.UpdatedAt,
		})
	}
	fmt.Println("All messages processed for chat ID:", chatID)

	return response.GetMessagesResponse{Messages: responseMessages}, nil
}
func MarkMessageRead(messageID string) (map[string]interface{}, error) {
	// TODO: Implement mark message read logic
	return map[string]interface{}{"message_id": messageID}, nil
}

func SendMessageOneToOne(receiverMobileNumber string, message string) (string, error) {
	// TODO: Implement send one-to-one message logic
	_, err := thirdparty.SendMessageRequest(receiverMobileNumber, message)

	if err != nil {
		return "Not able to send msg", err
	}

	return "Message sent successfully", nil
}

func getCustomerID(mobileNumber string) string {
	// TODO: Implement logic to retrieve customer ID based on mobile number
	var customer dbModels.Customer
	result := config.DB.Where("customers.mobile_number = ?", mobileNumber).First(&customer)
	if result.Error != nil {
		return ""
	}
	return customer.CustomerId
}
