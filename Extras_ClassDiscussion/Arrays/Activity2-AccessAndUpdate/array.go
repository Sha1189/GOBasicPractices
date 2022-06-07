package main

import "fmt"

func main() {
	var OperatingsystemList [6]string

	OperatingsystemList[0] = "Windwos XP"
	OperatingsystemList[1] = "Linux 1.0"
	OperatingsystemList[2] = "raspbian teensy"
	OperatingsystemList[3] = "MacOS"
	OperatingsystemList[4] = "IOS"
	OperatingsystemList[5] = "Google Android"

	OperatingsystemList[0] = "Windwos 10"
	OperatingsystemList[1] = "Linux 16.04"
	OperatingsystemList[2] = "Raspbian Buster"

	for i := range OperatingsystemList {
		fmt.Println(OperatingsystemList[i])
	}
}
