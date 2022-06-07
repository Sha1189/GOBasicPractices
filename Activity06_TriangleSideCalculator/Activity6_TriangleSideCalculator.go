package main

import (
	"fmt"
	"math"
)

func main() {
	var length1 float64
	var length2 float64
	var angle float64

	fmt.Println("Enter the first length of the triangle: ")
	fmt.Scanln(&length1)
	fmt.Println("Enter the second length of the triangle: ")
	fmt.Scanln(&length2)
	fmt.Println("Enter the angle between the two lengths: ")
	fmt.Scanln(&angle)

	length3 := math.Sqrt((length1 * length1) + (length2 * length2) - (2 * length1 * length2 * math.Cos(angle)))

	fmt.Printf("The 3rd length of the %.2f", length3)
}
