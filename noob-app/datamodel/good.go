package datamodel

type OrderGood struct {
	description string
	price       float64
	count       int
}

// func (orderGood *OrderGood) NewGood(
// 	desc string,
// 	price float64,
// 	count int) OrderGood {

// 	newGood := new(OrderGood)
// 	newGood.description = desc
// 	newGood.price = price
// 	newGood.count = count

// 	return *newGood
// }

func (orderGood *OrderGood) String() string {

	orderGoodStr := orderGood.description

	return orderGoodStr

}
