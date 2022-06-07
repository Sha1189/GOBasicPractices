package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var number string
		fmt.Println("Enter the number: ")
		fmt.Scanln(&number)

		numberValue, _ := strconv.ParseInt(number, 10, 0)

		if numberValue%2 == 0 {
			fmt.Println("The number is even!")
		} else {
			fmt.Println("The number is odd!")
		}
	}
}
