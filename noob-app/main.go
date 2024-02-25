package main

import (
	"fmt"
	"noob-app/datamodel"
)

func main() {

	newCustomer := datamodel.NewCustomer("Никита", "Буянов", "8-800-5553535")

	myOrder := datamodel.Order{
		ID:       1,
		Customer: newCustomer,
	}

	myOrder.AddGood("ak-74m", 1000.0, 1)
	myOrder.AddGood("5,45x39 Игольник", 499.0, 120)

	fmt.Println(myOrder.String())

}
