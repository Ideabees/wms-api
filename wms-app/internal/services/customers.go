package services

import (
	"errors"
	"fmt"
	"wms-app/config"
	"wms-app/internal/models/dbModels"
	"wms-app/internal/models/request"
	"wms-app/internal/models/response"
	thirdparty "wms-app/internal/third-party"
)

func CreateCustomer(customer *dbModels.Customer) (string, error) {

	// call third-party service to send opt-in message
	_, err := thirdparty.MakeOptInRequest(customer.MobileNumber)
	if err != nil {
		fmt.Println("Error in sending opt-in request", err)
		return "Failed to send opt-in message", err
	}
	// Insert customer into the database

	result := config.DB.Create(customer)
	if result.Error != nil {
		fmt.Println("Error in create user repository")
		return "Failed", result.Error
	}

	return "Success", nil
}

func GetCustomers(userID string, firstName string, lastName string) ([]response.GetCustomer, string, error) {

	resp := []response.GetCustomer{}
	var customers []dbModels.Customer
	result := config.DB.Where("customers.user_id = ?", userID).Find(&customers)
	if result.Error != nil {
		return resp, "Authorization Denied", result.Error
	}

	for _, cust := range customers {
		var rsp response.GetCustomer
		rsp.CustomerId = cust.CustomerId
		rsp.FirstName = cust.FirstName
		rsp.LastName = cust.LastName
		rsp.MobileNumber = cust.MobileNumber
		rsp.EmailID = cust.EmailID
		rsp.City = cust.City
		rsp.Pincode = cust.Pincode
		rsp.UpdatedOn = cust.UpdatedAt.String()
		rsp.CreatedBy = firstName + " " + lastName
		resp = append(resp, rsp)
	}
	// fetch user name from user id

	return resp, "Success", nil
}

func DeleteCustomers(customerIds *request.DeleteCustomer) (string, error) {

	// TODO: Define all usecases and erros in case of multiple customer ids
	for _, id := range customerIds.CustomerIds {
		fmt.Println("id")
		result := config.DB.Where("customers.customer_id = ?", id).Delete(&dbModels.Customer{})
		if result.Error != nil {
			fmt.Println("Error in delete user repository")
			return "Failed", result.Error
		}
		if result.RowsAffected == 0 {
			fmt.Println("Error in create user repository")
			return "No user found with this id", errors.New("no matching record found")
		}
	}
	return "Success", nil
}
