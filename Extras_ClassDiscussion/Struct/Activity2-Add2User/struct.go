package main

import "fmt"

type customer struct {
	fName            string
	lName            string
	age              int
	subscriber       bool
	homeAddress      string
	phone            int
	creditAvailable  float32
	currentCartCost  float32
	currentOrdercost float32
}

func main() {

	customer1 := customer{
		fName:            "Annakin",
		lName:            "Skywalker",
		age:              45,
		subscriber:       true,
		homeAddress:      "Death Star",
		phone:            1234567,
		creditAvailable:  10000.00,
		currentCartCost:  0.00,
		currentOrdercost: 0.00,
	}

	customer2 := customer{
		fName:            "Han",
		lName:            "Solo",
		age:              50,
		subscriber:       false,
		homeAddress:      "Tatooine",
		phone:            4321765,
		creditAvailable:  20000.00,
		currentCartCost:  0.00,
		currentOrdercost: 0.00,
	}

	fmt.Println(customer1.age)
	fmt.Println(customer2)
}
