package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	checkChar := "AZYXUTSRPMLKJHGEDCB"

	var carPlate string
	fmt.Println("Enter the car plate: ")
	fmt.Scanln(&carPlate)

	countLetter := 0
	countNonLetter := 0

	for val, r := range carPlate {
		fmt.Println(val, " ", r, " ", string(r))
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			countLetter++
		} else {
			countNonLetter++
		}

	}
	var total int
	if countLetter == 1 {
		total += (strings.Index(letters, string(carPlate[0])) + 1) * 4
	} else if countLetter == 2 {
		total += (strings.Index(letters, string(carPlate[0])) + 1) * 4
		total += (strings.Index(letters, string(carPlate[1])) + 1) * 9
	} else {
		total += (strings.Index(letters, string(carPlate[1])) + 1) * 9
		total += (strings.Index(letters, string(carPlate[2])) + 1) * 4
	}

	if countNonLetter == 4 {
		value1, _ := strconv.ParseInt(string(carPlate[len(carPlate)-1]), 10, 0)
		value2, _ := strconv.ParseInt(string(carPlate[len(carPlate)-2]), 10, 0)
		value3, _ := strconv.ParseInt(string(carPlate[len(carPlate)-3]), 10, 0)
		value4, _ := strconv.ParseInt(string(carPlate[len(carPlate)-4]), 10, 0)

		total += int(value1*2 + value2*3 + value3*4 + value4*5)

	} else if countNonLetter == 3 {
		value1, _ := strconv.ParseInt(string(carPlate[len(carPlate)-1]), 10, 0)
		value2, _ := strconv.ParseInt(string(carPlate[len(carPlate)-2]), 10, 0)
		value3, _ := strconv.ParseInt(string(carPlate[len(carPlate)-3]), 10, 0)

		total += int(value1*2 + value2*3 + value3*4)
	} else if countNonLetter == 2 {
		value1, _ := strconv.ParseInt(string(carPlate[len(carPlate)-1]), 10, 0)
		value2, _ := strconv.ParseInt(string(carPlate[len(carPlate)-2]), 10, 0)

		total += int(value1*2 + value2*3)
	} else {
		value1, _ := strconv.ParseInt(string(carPlate[len(carPlate)-1]), 10, 0)

		total += int(value1 * 2)
	}

	checkCharVal := total % 9

	fmt.Println("The checksum is ", string(checkChar[checkCharVal]))

}
