package main

import (
	"fmt"
)

func main() {
	var temperatureChoice int
	fmt.Println("Enter your temperature format (1 for Kelvin, 2 for Celsius, 3 for Fahrenheit): ")
	fmt.Scanln(&temperatureChoice)

	var temperatureInput float32
	fmt.Println("Enter the temperature: ")
	fmt.Scanln(&temperatureInput)

	if temperatureChoice == 1 {
		resultFahrenheit := (temperatureInput * (9.0 / 5.0)) - 459.67
		resultCelsius := (5.0 / 9.0) * (resultFahrenheit - 32)
		fmt.Println("Fahrenheit :", resultFahrenheit, " Celsius:", resultCelsius)
		fmt.Printf("Fahrenheit : %v %T\n", resultFahrenheit, resultCelsius)
	}
	if temperatureChoice == 2 {
		resultFahrenheit := temperatureInput*(9.0/5.0) + 32
		resultKelvin := (5.0 / 9.0) * (resultFahrenheit + 459.67)
		fmt.Println("Fahrenheit: ", resultFahrenheit, " Kelvin: ", resultKelvin)
	}
	if temperatureChoice == 3 {
		resultKelvin := (5.0 / 9.0) * (temperatureInput + 459.67)
		resultCelsius := (5.0 / 9.0) * (temperatureInput - 32)
		fmt.Println("Kelvin : ", resultKelvin, " Celsius: ", resultCelsius)
	}
}
