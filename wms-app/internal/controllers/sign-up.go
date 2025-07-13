package controllers

import (
	"net/http"
	"wms-app/internal/models/request"
	"wms-app/internal/models/dbModels"
	"wms-app/internal/services"
	"wms-app/internal/utils"

	"github.com/gin-gonic/gin"
)


// @Summary Register a new user
// @Description Registers a user with the provided details
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body RegisterRequest true "User registration data"
// @Success 200 {object} RegisterResponse
// @Router /api/register [post]
func Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
		return
	}
	// Create UUID
	id := utils.CreateUUID()

	user := dbModels.User{
		UserId: id,
		FirstName:    req.FirstName,
		MiddleName:   req.MiddleName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: hash,
		MobileNumber: req.MobileNumber,
	}

	response, err := services.CreateUser(&user)
	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
		return
	}

	token, err := utils.GenerateToken(user.UserId, user.Email, user.FirstName, user.LastName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data" : response,
		"token":   token,
	})
}
