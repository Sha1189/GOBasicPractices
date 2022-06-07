package main

import (
	"fmt"
	"math"
)

/*
func calSqrt(a float64) (result float64, ok bool) {
	if a >= 0 {
		result = math.Sqrt(a)
		ok = true
	}
	return result, ok
}
*/

func calSqrt(a, b float64) (result float64, ok bool) {
	if a < 0 {
		return
	}
	return math.Sqrt(a), true
}

func main() {
	res, err := calSqrt(2, 2)
	fmt.Println(res, err)
}
