package main

import "fmt"

func main() {
	var customer1 = customer{FName: "Micheal", LName: "Jordan", Username: "MJ2020", Password: "1234567",
		Email: "MJ2020@gmail.com", Phone: 12345678, Address: "18227 Capstan Greens Road Cornelius, NC 28031"}

	fmt.Println(customer1)
	customer1.printAllInfo()
	fmt.Println(customer1.printUserCredentials())
	fmt.Println(customer1.printUserAddress())

}
