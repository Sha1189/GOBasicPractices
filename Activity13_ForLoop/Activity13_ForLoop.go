package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 1000; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 1000; i++ {
		if i%2 != 0 {
			fmt.Println("Odd", i)
		} else {
			fmt.Println("Even", i)
		}
	}
	for counter := 0; counter <= 1000; counter++ {
		if counter%2 == 0 {
			fmt.Println("Even", counter)
		}
	}

}
