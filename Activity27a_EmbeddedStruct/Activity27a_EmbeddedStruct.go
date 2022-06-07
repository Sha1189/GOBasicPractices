package main

import "fmt"

type customer struct {
	customerName    string
	customerID      int
	customerAddress string
}

type orderItems struct {
	itemName     string
	itemQuantity int
	itemCost     float64
}

type order struct {
	customer
	orderNumber int
	orderItems
}

func main() {
	order1 := order{
		customer: customer{
			customerName:    "Anakin Skywalker",
			customerID:      123456,
			customerAddress: "Death Star",
		},
		orderNumber: 123456765,
		orderItems: orderItems{
			itemName:     "DeathStar",
			itemQuantity: 1,
			itemCost:     1000000000,
		},
	}

	order2 := struct {
		annoymousInfo  int
		annoymousInfo2 int
		customer
		detail string
	}{annoymousInfo: 1213445,
		annoymousInfo2: 123123,
		customer: customer{
			customerName:    "Anonymous Customer",
			customerAddress: "Anonymous Address",
		},
		detail: "new field"
		}

	fmt.Println(order1.customerID)
	fmt.Println(order1.customer.customerID)
	fmt.Println(order2)
}
