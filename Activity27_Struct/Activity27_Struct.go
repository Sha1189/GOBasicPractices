package main

import "fmt"

type customerTable struct{
	id int
	customerDetail customer
}

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

	member1 := customer{
		id int}{id: 12345},

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

	customerbase := []customer{customer1, customer2}
	fmt.Println(member1., customer2.fName)

	for _, res := range customerbase {
		fmt.Println(res.fName)
	}

}
