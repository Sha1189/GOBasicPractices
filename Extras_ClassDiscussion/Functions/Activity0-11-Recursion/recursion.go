package main

import "fmt"

// binary tree/ go advanced/ brank, keep searching
//passing argument/ simple coundown use for loop

func countdown(counter int) {
	if counter == 0 {
		fmt.Println("0")
		return
	} else {
		fmt.Println(counter)
		countdown(counter - 1)
	}
}

func main() {
	countdown(5)
}
