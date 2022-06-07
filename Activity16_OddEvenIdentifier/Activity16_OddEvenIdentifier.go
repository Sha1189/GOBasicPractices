package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var number string
		fmt.Println("Enter the first number: ")
		fmt.Scanln(&number)

		numberValue, _ := strconv.ParseInt(number, 10, 0)
		if numberValue%2 == 0 {
			fmt.Println("It is EVEN!")
		} else {
			fmt.Println("It is ODD!")
		}
		if numberValue >= 10 {
			fmt.Println("It has more than 1 digit!")
		} else {
			fmt.Println("It has only 1 digit!")
		}
	}
}
