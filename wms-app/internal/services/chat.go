package services

import (
	thirdparty "wms-app/internal/third-party"
)

func CreateChat(chatData map[string]interface{}) (map[string]interface{}, error) {
	// TODO: Implement chat creation logic
	return chatData, nil
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
