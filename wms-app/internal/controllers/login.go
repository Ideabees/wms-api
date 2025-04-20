package controllers

import (
	"net/http"
	"wms-app/config"
	"wms-app/internal/models/dbModels"
	"wms-app/internal/models/request"
	"wms-app/internal/models/response"
	"wms-app/internal/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user dbModels.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// build response
	var response response.CreateUserRespone
	response.UserId = user.UserId
	response.FirstName = user.FirstName
	response.MiddleName = user.MiddleName
	response.LastName = user.LastName
	response.Email = user.Email
	response.MobileNumber = user.MobileNumber
	response.CreatedAt = user.CreatedAt.String()
	response.UpdatedAt = user.UpdatedAt.String()

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data": response,
		"token":   token,
	})
}

func Logout(c *gin.Context) {
	token := c.GetString("token_string")
	utils.BlacklistToken(token)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

