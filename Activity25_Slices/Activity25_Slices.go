package main

import "fmt"

func main() {
	spending := []float64{9.5, 8, 10.2, 7.4}

	fmt.Println(len(spending))
	fmt.Println(cap(spending))

	spending[2] = 9.8

	fmt.Println(spending)

	spending = append(spending, 8.4, 9.4, 7.2)
	fmt.Println(len(spending))
	fmt.Println(cap(spending))

	subslice := spending[3:]
	fmt.Println(subslice)
	fmt.Println(len(subslice))
	fmt.Println(cap(subslice))
}
