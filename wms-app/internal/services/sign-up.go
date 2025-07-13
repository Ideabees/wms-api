package services

import (
	"fmt"
	"wms-app/config"
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

func CreateUserWithDB(db DBInterface, user *dbModels.User) (response.CreateUserRespone, error) {
	var response response.CreateUserRespone
	result := db.Create(user)
	if result.Error() != nil {
		fmt.Println("Error in create user repository")
		return response, result.Error()
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

// Backward compatible wrapper
func CreateUser(user *dbModels.User) (response.CreateUserRespone, error) {
	return CreateUserWithDB(GormDB{db: config.DB}, user)
}
