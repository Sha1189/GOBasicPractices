package main

import "fmt"

type customer struct {
	fName    string
	lName    string
	username string
	password string
	email    string
	phone    int
	address  string
}

var customer1 customer

func (c customer) printUserCredentials() (string, string) {
	return c.username, c.password
}

func (c customer) printUserAddress() string {
	return c.address
}

func (c customer) printAllInfo() {
	fmt.Println(c.fName, c.lName, c.username, c.password, c.email, c.phone, c.address)
}
