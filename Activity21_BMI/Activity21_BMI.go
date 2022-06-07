package main

import (
	"fmt"
	"strconv"
)

func main() {
	var height string
	var weight string
	fmt.Println("Enter your height: ")
	fmt.Scanln(&height)
	fmt.Println("Enter your weight: ")
	fmt.Scanln(&weight)

	heightValue, _ := strconv.ParseFloat(height, 64)
	weightValue, _ := strconv.ParseFloat(weight, 64)

	switch BMI := weightValue / (heightValue * heightValue); {
	case BMI < 18.5:
		fmt.Println("Underweight")
	case BMI >= 18.5 && BMI < 24.9:
		fmt.Println("Healthy Weight")
	case BMI >= 25 && BMI < 29.9:
		fmt.Println("Overweight")
	case BMI >= 30 && BMI < 34.9:
		fmt.Println("Obese")
	case BMI >= 35 && BMI < 39.9:
		fmt.Println("Severely Obese")
	default:
		fmt.Println("Morbidly Obese")
	}
}
