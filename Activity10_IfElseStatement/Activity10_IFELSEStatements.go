package main

import (
	"fmt"
	"strconv"
)

func main() {
	var compare string
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&compare)
	fmt.Printf("Type is %T\n", compare)

	compareValue, _ := strconv.ParseInt(compare, 10, 0)
	/*	if err != nil {
			fmt.Println("Error in conversion")
			log.Fatal("Error.... stopping program!")
		}
	*/
	if compareValue%5 == 0 && compareValue%6 == 0 {
		fmt.Println("The value is divisible by 5 and 6 !")
	} else {
		fmt.Println("The value is NOT divisible by 5 and 6 !")
	}
}
