package main

import (
	"fmt"
)

func main() {
	if number := 17; number%2 == 0 {
		fmt.Println("It is EVEN!")
	} else {
		fmt.Println("It is ODD!")
	}
	if number := 17; number >= 10 {
		fmt.Println("It has more than 1 digit!")
	} else {
		fmt.Println("It has only 1 digit!")
	}
}
