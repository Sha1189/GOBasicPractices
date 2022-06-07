package main

import "fmt"

func main() {

	var OperatingsystemList [6]string

	OperatingsystemList[0] = "Windows XP"
	OperatingsystemList[1] = "Linux 1.0"
	OperatingsystemList[2] = "raspbian teensy"
	OperatingsystemList[3] = "MacOS"
	OperatingsystemList[4] = "IOS"
	OperatingsystemList[5] = "Google Android"

	for i := range OperatingsystemList {
		fmt.Println(len(OperatingsystemList[i]))
	}

}
