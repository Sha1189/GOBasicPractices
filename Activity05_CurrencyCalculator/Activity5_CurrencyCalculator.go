package main

import (
	"fmt"
	"math"
)

func main() {
	var oneDollarCoin float32
	var fiftyCentCoin float32
	var twentyCentCoin float32
	var tenCentCoin float32
	var fiveCentCoin float32

	/*	var a float32
		var b float32
		var c float32
		var d float32
		var e float32
	*/
	fmt.Println("Enter number of 1-dollar coins: ")
	fmt.Scanln(&oneDollarCoin)
	fmt.Println("Enter number of 50-cents coins: ")
	fmt.Scanln(&fiftyCentCoin)
	fmt.Println("Enter number of 20-cents coins: ")
	fmt.Scanln(&twentyCentCoin)
	fmt.Println("Enter number of 10-cents coins: ")
	fmt.Scanln(&tenCentCoin)
	fmt.Println("Enter number of 5-cents coins: ")
	fmt.Scanln(&fiveCentCoin)

	totalAmount := oneDollarCoin + fiftyCentCoin*0.5 + twentyCentCoin*0.2 + tenCentCoin*0.1 + fiveCentCoin*0.05
	fmt.Println("Total amount:", totalAmount)
	fmt.Println("Number of 2-dollar notes: ", math.Floor(float64(totalAmount/2)))
	fmt.Println("Amount of change: ", math.Mod(float64(totalAmount), float64(2)))
}
