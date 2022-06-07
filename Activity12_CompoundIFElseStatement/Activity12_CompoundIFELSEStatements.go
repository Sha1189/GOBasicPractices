package main

import (
	"fmt"
	"strconv"
)

func main() {
	var year string
	fmt.Println("Enter the year: ")
	fmt.Scanln(&year)

	yearInt, _ := strconv.ParseInt(year, 10, 0)

	if (yearInt%4 == 0 && yearInt%100 != 0) || yearInt%400 == 0 {
		fmt.Println(yearInt, "is a leap year.")
	} else {
		fmt.Println(yearInt, "is not a leap year.")
	}
}
