package main

import "fmt"

func main() {
	number1 := 30
	number2 := 30

	if number1 > number2 {
		difference := number1 - number2
		fmt.Println(number1, "is bigger!")
		fmt.Println("The difference is ", difference)
	}
	if number1 < number2 {
		difference := number2 - number1
		fmt.Println(number2, "is bigger!")
		fmt.Println("The difference is ", difference)
	}

}
