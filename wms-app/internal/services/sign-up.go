package services

import (
	"fmt"
	"wms-app/config"
	"wms-app/internal/models/dbModels"
	"wms-app/internal/models/response"
)


func CreateUser(user *dbModels.User) (response.CreateUserRespone, error) {

	var response response.CreateUserRespone
	result := config.DB.Create(user)
	if result.Error != nil {
		fmt.Println("Error in create user repository")
		return response, result.Error
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

	return response, result.Error
}
