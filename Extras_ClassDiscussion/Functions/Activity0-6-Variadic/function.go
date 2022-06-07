package main

import "fmt"

func main() {
	//fmt.Println("First Total", sum())
	//fmt.Println("Second total", sum(3))
	fmt.Println("last total", sum(1, 2, 3, 6, 4, 3, 2, 1))

}

func sum(vals ...int) int {
	total := 0
	fmt.Println(vals)
	for i, val := range vals {
		fmt.Println("index", i, "Value", val)
		total += val
	}
	return total
}
