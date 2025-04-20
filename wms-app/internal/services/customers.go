package services

import (
	"fmt"
	"wms-app/config"
	"wms-app/internal/models/dbModels"
)

func CreateCustomers(customer *dbModels.Customer) (string, error) {

	result := config.DB.Create(customer)
	if result.Error != nil {
		fmt.Println("Error in create user repository")
		return "Failed", result.Error
	}

	return "Success", result.Error
}
