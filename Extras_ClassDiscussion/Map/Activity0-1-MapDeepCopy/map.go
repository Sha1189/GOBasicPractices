package main

import "fmt"

func main() {
	originalMap := make(map[string]int)
	originalMap["Vegetables"] = 1
	originalMap["Drinks"] = 2
	originalMap["Meat"] = 3
	originalMap["Appliances"] = 4

	targetMap := make(map[string]int)

	for key, value := range originalMap {
		targetMap[key] = value
	}

	fmt.Println(targetMap)

}
