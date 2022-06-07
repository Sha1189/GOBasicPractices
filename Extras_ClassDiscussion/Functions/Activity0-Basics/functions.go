package main

import "fmt"

var a int = 20
var b int = 50

func main() {
	fmt.Println(a)
	fmt.Println(sum(a, b))
	a = 50
	fmt.Println(a)
	fmt.Println(sum(a, b))
	fmt.Println(a)
}

func sum(a, b int) int {
	c := a + b
	return c
}
