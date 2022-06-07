package main

import "fmt"

func shippingCost(length, width, height float64) (float64, float64) {
	var airCost = length * width * height * 1.25
	var shippingCost = length * width * height * 0.75
	return airCost, shippingCost
}

func main() {
	costByAir, _ := shippingCost(10.5, 12.6, 17.5)
	fmt.Printf("Cost by Air is $%.2f", costByAir)
}
