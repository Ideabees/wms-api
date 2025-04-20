package request

type CreateCustomersRequest struct {
	FirstName       string `json:"first_name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	MobileNumber    string `json:"mobile_number" binding:"required"`
}