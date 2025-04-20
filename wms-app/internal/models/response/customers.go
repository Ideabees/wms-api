package response


type GetCustomer struct {
	FirstName    string
	LastName     string
	MobileNumber string
	CreatedBy    string
	UpdatedOn    string
}

type GetCustomers struct {
	Customers []GetCustomer
}