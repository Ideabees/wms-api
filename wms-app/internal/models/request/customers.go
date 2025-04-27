package request

type CreateCustomerRequest struct {
	FirstName       string `json:"first_name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	MobileNumber    string `json:"mobile_number" binding:"required"`
}

type DeleteCustomer struct {
	CustomerIds     []string `json:"customer_ids" binding:"required"`
}