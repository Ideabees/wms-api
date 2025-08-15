package services

import (
	"fmt"
	"wms-app/config"
	"wms-app/internal/models/dbModels"
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
func GetUserChats(userID string) ([]map[string]interface{}, error) {
	// TODO: Implement user chats retrieval logic
	return []map[string]interface{}{{"user_id": userID}}, nil
}
func SendMessage(chatID string, messageData map[string]interface{}) (map[string]interface{}, error) {
	// TODO: Implement send message logic
	return messageData, nil
}
func GetMessages(chatID string) ([]map[string]interface{}, error) {
	// TODO: Implement get messages logic
	return []map[string]interface{}{{"chat_id": chatID}}, nil
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