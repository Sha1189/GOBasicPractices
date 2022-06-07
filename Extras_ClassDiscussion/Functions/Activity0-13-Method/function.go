package main

import "fmt"

type totalCost float64

func (tc totalCost) Total() float64 {
	tc = tc + tc*0.07
	return float64(tc)
}

func main() {
	var tc totalCost = 1
	fmt.Println(tc.Total())
	fmt.Println(tc.Total())
}
