package main

import (
	"fmt"
)

func main() {
	var firstValue float32
	fmt.Println("Enter an integer value: ")
	fmt.Scanln(&firstValue)

	var mathFunction string
	fmt.Println("Enter a mathematical operation: ")
	fmt.Scanln(&mathFunction)

	var secondValue float32
	fmt.Println("Enter an integer value: ")
	fmt.Scanln(&secondValue)

	var result float32
	if plus := "+"; mathFunction == plus {
		result = firstValue + secondValue
	}
	if minus := "-"; mathFunction == minus {
		result = firstValue - secondValue
	}
	if multiply := "*"; mathFunction == multiply {
		result = firstValue * secondValue
	}
	if divide := "/"; mathFunction == divide {
		result = firstValue / secondValue
	}
	fmt.Println(result)
}
