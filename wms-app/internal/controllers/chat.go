package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateChat handles the creation of a new chat
func CreateChat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Chat created successfully"})
}

// GetChat retrieves a specific chat by ID
func GetChat(c *gin.Context) {
	chatID := c.Param("chat_id")
	c.JSON(http.StatusOK, gin.H{"message": "Chat retrieved successfully", "chat_id": chatID})
}

// GetUserChats retrieves all chats for a specific user
func GetUserChats(c *gin.Context) {
	userID := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{"message": "User chats retrieved successfully", "user_id": userID})
}

// SendMessage handles sending a message to a chat
func SendMessage(c *gin.Context) {
	chatID := c.Param("chat_id")
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully", "chat_id": chatID})
}

// GetMessages retrieves all messages for a specific chat
func GetMessages(c *gin.Context) {
	chatID := c.Param("chat_id")
	c.JSON(http.StatusOK, gin.H{"message": "Messages retrieved successfully", "chat_id": chatID})
}

// MarkMessageRead marks a specific message as read
func MarkMessageRead(c *gin.Context) {
	messageID := c.Param("message_id")
	c.JSON(http.StatusOK, gin.H{"message": "Message marked as read", "message_id": messageID})
}
