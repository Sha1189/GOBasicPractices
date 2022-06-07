package main

import "fmt"

//func (recevier) identifier ( ?? parameters) (returns)
func sum(a, b int) {
	result := a + b
	fmt.Println(result)
}

func main() {
	fmt.Println("This is the first in")
	defer fmt.Println("This is the last in")
	defer sum(10, 20)
	defer sum(1, 1)
}

// after function is done then you call out defer
// execute defer after all work is done
