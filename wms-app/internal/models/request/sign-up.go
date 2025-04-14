package request

type RegisterRequest struct {
	FirstName       string `json:"first_name" binding:"required"`
	MiddleName      string `json:"middle_name"`
	LastName        string `json:"last_name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	MobileNumber    string `json:"mobile_number" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}