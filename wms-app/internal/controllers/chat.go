package controllers

import (
	"net/http"
	"time"
	"wms-app/internal/models/dbModels"
	"wms-app/internal/models/request"
	"wms-app/internal/services"
	"wms-app/internal/utils"

	"github.com/gin-gonic/gin"
)

// CreateChat handles the creation of a new chat
func CreateChat(c *gin.Context) {
	// TODO: Implement chat creation logic

	chatRequest := request.CreateChat{}
	if err := c.ShouldBindJSON(&chatRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chat_id := utils.CreateUUID()

	chat := dbModels.CreateChats{
		ChatID:               chat_id,
		SenderID:             c.GetString("user_id"),
		ReceiverID:           "",
		ReceiverMobileNumber: chatRequest.ReceiverMobileNumber,
		CreatedAt:           time.Now(),
	}

	// Call the service to create the chat
	id, err := services.CreateChat(&chat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not able to create chat", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chat created successfully", "chat_id": id})
}

// GetChat retrieves a specific chat by ID
func GetChat(c *gin.Context) {
	chatID := c.Param("chat_id")
	c.JSON(http.StatusOK, gin.H{"message": "Chat retrieved successfully", "chat_id": chatID})
}

// GetUserChats retrieves all chats for a specific user
func GetUserChats(c *gin.Context) {
	userID := c.Param("user_id")
	data, err := services.GetUserChats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not able to retrieve user chats", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User chats retrieved successfully", "user_id": userID, "data": data})
}

// SendMessage handles sending a message to a chat
func SendMessage(c *gin.Context) {
	var messageData request.SendMessageOneToOne
	if err := c.ShouldBindJSON(&messageData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to send the message

	// build db models mapper

	msg_id := utils.CreateUUID()
	var msg dbModels.CreateMessages

	if messageData.MessageType == "text" {
		// Handle text message specific logic
		msg.MessageID = msg_id
		msg.ChatID = messageData.ChatID
		msg.SenderID = c.GetString("user_id")
		msg.MessageType = messageData.MessageType
		msg.Content = messageData.Message
		msg.MediaURL = ""
		msg.IsReadByReceiver = "false"
		msg.CreatedAt = time.Now()
		msg.UpdatedAt = time.Now()
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "send only text type"})
	}

	_, err := services.SendMessage(&msg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not able to send message", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully", "msg_id": msg.MessageID})
}

// GetMessages retrieves all messages for a specific chat
func GetMessages(c *gin.Context) {
	chatID := c.Param("chat_id")
	data, err := services.GetMessages(chatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not able to retrieve messages", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Messages retrieved successfully", "chat_id": chatID, "data": data})
}

// MarkMessageRead marks a specific message as read
func MarkMessageRead(c *gin.Context) {
	messageID := c.Param("message_id")
	c.JSON(http.StatusOK, gin.H{"message": "Message marked as read", "message_id": messageID})
}

// SendMessageOneToOne handles sending a one-to-one message
func SendMessageOneToOne(c *gin.Context) {
	// TODO: Implement send one-to-one message logic
	var req request.SendMessageOneToOne
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receiverMobileNumber := req.ReceiverMobileNumber
	message := req.Message
	_, err := services.SendMessageOneToOne(receiverMobileNumber, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not able to send message", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}
