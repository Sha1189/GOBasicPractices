package main

import "fmt"

// declare this when you need to declare a function /struc
// slice of function
type newFunction func(int, int) int

var a = 1
var b = 2

func main() {
	funcSet := []newFunction{
		func(a int, b int) int { return a + b },
		func(a int, b int) int { return a - b },
		func(a int, b int) int { return a * b },
		func(a int, b int) int { return a / b },
	}

	var fn = funcSet[0]
	fmt.Println("a+b = ", fn(a, b))
}
