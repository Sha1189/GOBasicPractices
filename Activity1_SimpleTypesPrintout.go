package main

import (
	"fmt"
)

func main() {
	var text string = "The following is the account information"
	firstName := "Luke"
	secondName := "Skywalker"
	age := 20
	weight := 73.0
	height := 1.72
	remainingCredit := 123.55
	accountName := "admin"
	accountPassword := "password"
	subscribed := true

	fmt.Printf("%v -  %T\n", text, text)
	fmt.Printf("%v -  %T\n", firstName, firstName)
	fmt.Printf("%v -  %T\n", secondName, secondName)
	fmt.Printf("%v -  %T\n", age, age)
	fmt.Printf("%v -  %T\n", weight, weight)
	fmt.Printf("%v -  %T\n", height, height)
	fmt.Printf("%v -  %T\n", remainingCredit, remainingCredit)
	fmt.Printf("%v -  %T\n", accountName, accountName)
	fmt.Printf("%v -  %T\n", accountPassword, accountPassword)
	fmt.Printf("%v -  %T\n", subscribed, subscribed)
}
