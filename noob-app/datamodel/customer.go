package datamodel

type OrderCustomer struct {
	name        string
	lastName    string
	phoneNumber string
}

func NewCustomer(name string, lastName string, phoneNum string) *OrderCustomer {
	newCustomer := OrderCustomer{name: name, lastName: lastName, phoneNumber: phoneNum}
	return &newCustomer
}
