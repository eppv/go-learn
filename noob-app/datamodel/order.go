package datamodel

import "noob-app/utils"

type Order struct {
	ID       int
	Customer *OrderCustomer
	Goods    []*OrderGood
}

func (order *Order) String() string {

	// var orderGoods []string

	var orderGoodStr []string
	for _, v := range order.Goods {
		orderGoodStr = append(orderGoodStr, v.String())
	}

	orderStrSlice := []string{
		"Order(",
		"goods=[",
		utils.Concat(", ", orderGoodStr...),
		"]; ",
		"customer=",
		order.Customer.name + " " + order.Customer.lastName,
		"; ",
		"contact_number=",
		order.Customer.phoneNumber,
		")",
	}

	orderStr := utils.Concat("", orderStrSlice...)

	return orderStr
}

func (order *Order) AddGood(desc string, price float64, count int) {

	newGood := new(OrderGood)
	newGood.description = desc
	newGood.price = price
	newGood.count = count

	order.Goods = append(order.Goods, newGood)
}
