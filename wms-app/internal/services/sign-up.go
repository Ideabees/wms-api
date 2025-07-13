package services

import (
	"wms-app/config"
	"wms-app/internal/errors"
	"wms-app/internal/models/dbModels"
	"wms-app/internal/models/response"

	"gorm.io/gorm"
)

type DBInterface interface {
	Create(value interface{}) DBResult
}

type DBResult interface {
	Error() error
}

type GormDBResult struct {
	result *gorm.DB
}

func (g GormDBResult) Error() error {
	return g.result.Error
}

type GormDB struct {
	db *gorm.DB
}

func (g GormDB) Create(value interface{}) DBResult {
	return GormDBResult{result: g.db.Create(value)}
}

func CreateUserWithDB(db DBInterface, user *dbModels.User) (response.CreateUserRespone, *errors.ErrorResponse) {
	var response response.CreateUserRespone

	// Validate user input
	if user.Email == "" {
		return response, &errors.ErrorResponse{
			StatusCode: errors.ErrBadRequest,
			Error:      "ValidationError",
			Message:    "Email is required",
		}
	}

	if user.FirstName == "" {
		return response, &errors.ErrorResponse{
			StatusCode: errors.ErrBadRequest,
			Error:      "ValidationError",
			Message:    "First name is required",
		}
	}

	// Check for existing user with same email or mobile number
	var existingUser dbModels.User
	gormDB, ok := db.(GormDB)
	if !ok {
		return response, &errors.ErrorResponse{
			StatusCode: errors.ErrInternalServerError,
			Error:      "DatabaseError",
			Message:    "Invalid database type",
		}
	}
	checkResult := gormDB.db.Where("email = ? OR mobile_number = ?", user.Email, user.MobileNumber).First(&existingUser)
	if checkResult.Error == nil {
		return response, &errors.ErrorResponse{
			StatusCode: errors.ErrConflict,
			Error:      "DuplicateError",
			Message:    "User with this email or mobile number already exists",
		}
	}

	result := db.Create(user)
	if result.Error() != nil {
		if result.Error() == gorm.ErrDuplicatedKey {
			return response, &errors.ErrorResponse{
				StatusCode: errors.ErrConflict,
				Error:      "DuplicateError",
				Message:    "User with this email or mobile number already exists",
			}
		}
		return response, &errors.ErrorResponse{
			StatusCode: errors.ErrInternalServerError,
			Error:      "DatabaseError",
			Message:    "Failed to create user",
		}
	}

	// prep the response back
	response.UserId = user.UserId
	response.FirstName = user.FirstName
	response.MiddleName = user.MiddleName
	response.LastName = user.LastName
	response.Email = user.Email
	response.MobileNumber = user.MobileNumber
	response.CreatedAt = user.CreatedAt.String()
	response.UpdatedAt = user.UpdatedAt.String()
	return response, nil
}

// CreateUser creates a new user in the system
func CreateUser(user *dbModels.User) (response.CreateUserRespone, *errors.ErrorResponse) {
	return CreateUserWithDB(GormDB{db: config.DB}, user)
}
