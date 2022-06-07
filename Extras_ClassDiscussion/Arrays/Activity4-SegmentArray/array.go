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

	for i := 0; i < len(OperatingsystemList); i++ {
		NewOperatingSystemList[i] = OperatingsystemList[i]
	}

	NewOperatingSystemList[6] = "Ubuntu"
	NewOperatingSystemList[7] = "Ms-DOS"
	NewOperatingSystemList[8] = "Solaris"
	fmt.Println(OperatingsystemList)
	fmt.Println(NewOperatingSystemList)

	var i, j int

	for i = 0; i < 3; i++ {
		fmt.Printf("This is Set %v\n", i+1)
		for j = 0 + i*3; j < 3+i*3; j++ {
			fmt.Println(NewOperatingSystemList[j])
		}
		fmt.Println("------------")
	}
}
