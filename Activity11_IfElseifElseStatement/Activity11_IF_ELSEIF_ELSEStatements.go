package main

import (
	"fmt"
)

func main() {
	var username string
	fmt.Println("Enter your username: ")
	fmt.Scanln(&username)

	if username == "Robin" || username == "John" { //expression condition 1
		fmt.Println("Welcome!")
	} else if username == "Admin" { // condition 2
		fmt.Println("Welcome Admin!")
	} else {
		fmt.Println("Merry men")
	}
}
