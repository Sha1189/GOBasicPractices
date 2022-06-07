package main

import (
	"fmt"
	"strconv"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {

	var factorialVal string
	fmt.Println("Enter your factorial: ") //n* n-1 * n-2....
	fmt.Scanln(&factorialVal)

	factorialValue, _ := strconv.ParseInt(factorialVal, 10, 0)

	fmt.Println(factorial(int(factorialValue)))

}
