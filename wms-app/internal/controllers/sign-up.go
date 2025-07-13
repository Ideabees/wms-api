package controllers

import (
	"net/http"
	"wms-app/internal/errors"
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
		errResp := errors.NewErrorResponse(
			errors.ErrBadRequest,
			"ValidationError",
			err.Error(),
		)
		c.JSON(errResp.StatusCode, errResp)
		return
	}

	if req.Password != req.ConfirmPassword {
		errResp := errors.NewErrorResponse(
			errors.ErrBadRequest,
			"ValidationError",
			"Passwords do not match",
		)
		c.JSON(errResp.StatusCode, errResp)
		return
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		errResp := errors.NewErrorResponse(
			errors.ErrInternalServerError,
			"EncryptionError",
			"Password encryption failed",
		)
		c.JSON(errResp.StatusCode, errResp)
		return
	}

	// Create UUID
	id := utils.CreateUUID()
	
	// Create user model
	user := &dbModels.User{
		UserId:       id,
		FirstName:    req.FirstName,
		MiddleName:   req.MiddleName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: hash,
		MobileNumber: req.MobileNumber,
	}

	// Call service to create user
	resp, errResp := services.CreateUser(user)
	if errResp != nil {
		c.JSON(errResp.StatusCode, errResp)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(resp.UserId, resp.Email, resp.FirstName, resp.LastName)
	if err != nil {
		errResp := errors.NewErrorResponse(
			errors.ErrInternalServerError,
			"TokenError",
			"Failed to generate authentication token",
		)
		c.JSON(errResp.StatusCode, errResp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    resp,
		"message": "User registered successfully",
		"token":   token,
	})
}
