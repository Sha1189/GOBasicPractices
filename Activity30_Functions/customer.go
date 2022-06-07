package main

import "fmt"

type customer struct {
	FName    string
	LName    string
	Username string
	Password string
	Email    string
	Phone    int
	Address  string
}

func (c customer) printUserCredentials() (string, string) {
	fmt.Println(c)
	return c.Username, c.Password
}

func (c customer) printUserAddress() string {
	return c.Address
}

func (c customer) printAllInfo() {
	fmt.Println(c.FName, c.LName, c.Username, c.Password, c.Email, c.Phone, c.Address)
}
