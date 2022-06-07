package main

import (
	"fmt"
)

func main() {
	var OperatingSystemList [6]string

	OperatingSystemList[0] = "Windows XP" //length of 10
	OperatingSystemList[1] = "Linux 1.0"  //length of 9
	OperatingSystemList[2] = "Raspbian Teensy"
	OperatingSystemList[3] = "MacOS"
	OperatingSystemList[4] = "IOS"
	OperatingSystemList[5] = "Google Android"

	for i := 0; i < len(OperatingSystemList); i++ {
		fmt.Println(len(OperatingSystemList[i]))
	}

	OperatingSystemList[0] = "Windows 10"
	OperatingSystemList[1] = "Linux 16.04"
	OperatingSystemList[2] = "Raspbian Buster"

	fmt.Println(OperatingSystemList)

	var NewOperatingSystemList [9]string

	for i := 0; i < len(OperatingSystemList); i++ {
		NewOperatingSystemList[i] = OperatingSystemList[i]
	}
	NewOperatingSystemList[6] = "Ubuntu"
	NewOperatingSystemList[7] = "MS-Dos"
	NewOperatingSystemList[8] = "Solaris"

	fmt.Println(NewOperatingSystemList)

	fmt.Println(NewOperatingSystemList[0:3])
	fmt.Println(NewOperatingSystemList[3:6])
	fmt.Println(NewOperatingSystemList[6:])
}
