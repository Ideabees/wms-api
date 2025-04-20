package controllers

import (
	"fmt"
	"net/http"
	"wms-app/internal/models/dbModels"
	"wms-app/internal/models/request"
	"wms-app/internal/services"
	"wms-app/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	userID := c.GetString("user_id")
	//email := c.GetString("email")

	var req request.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cust_id := utils.CreateUUID()
	custModel := dbModels.Customer{
		CustomerId:   cust_id,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		MobileNumber: req.MobileNumber,
		UserId:       userID,
	}

	// call service layer to insert the customer
	msg, err := services.CreateCustomer(&custModel)
	if err != nil {
		fmt.Println("DB insertion Failed", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "DB insertion Failed",
			"status":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Customer succefully created",
			"status":  msg,
		})
	}
}

func GetCustomers(c *gin.Context) {
	userID := c.GetString("user_id")

	// call service layer to insert the customer
	data, msg, err := services.GetCustomers(userID)
	if err != nil {
		fmt.Println("DB insertion Failed", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "DB Operation Failed",
			"status":  msg,
			"data":    data,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "",
			"status":  msg,
			"data":    data,
		})
	}
}
