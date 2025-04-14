package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPermissions(c *gin.Context) {
	userID := c.GetString("user_id")
	email := c.GetString("email")

	c.JSON(http.StatusOK, gin.H{
		"message": "Protected profile access",
		"user_id": userID,
		"email":   email,
	})
}
