package main

import "fmt"

func main() {
	a := 12
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(&a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	*b = 10
	fmt.Println(a)
}
