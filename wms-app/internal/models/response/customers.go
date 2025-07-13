package response

type GetCustomer struct {
	CustomerId   string
	FirstName    string
	LastName     string
	MobileNumber string
	EmailID      string
	City         string
	Pincode      string
	CreatedBy    string
	UpdatedOn    string
}

type GetCustomers struct {
	Customers []GetCustomer
}
