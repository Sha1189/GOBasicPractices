package main

import "fmt"

func main() {
	stringA := "This is the message to print!"

	//count, err := fmt.Println(stringA)
	//fmt.Println(count, err)

	_, err := fmt.Println(stringA)
	fmt.Println(err)
}
