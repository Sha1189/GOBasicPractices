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

	var NewOperatingSystemList [9]string
	for i := range OperatingsystemList {
		NewOperatingSystemList[i] = OperatingsystemList[i]
	}
	NewOperatingSystemList[6] = "Ubuntu"
	NewOperatingSystemList[7] = "Ms-DOS"
	NewOperatingSystemList[8] = "Solaris"
	fmt.Println(OperatingsystemList)
	fmt.Println(NewOperatingSystemList)
}
