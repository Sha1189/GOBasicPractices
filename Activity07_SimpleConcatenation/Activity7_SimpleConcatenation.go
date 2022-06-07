package main

import (
	"fmt"
)

func main() {
	list1 := [4]string{"ans", "wer", "is", "of"}
	list2 := [4]string{"-", "+", "*", "/"}
	list3 := [4]string{"5", "10", "4", "9"}

	fmt.Printf("%s%s %s %s %s %s %s %s", list1[0], list1[1], list1[3], list3[0], list2[1], list3[2], list1[2], list3[3])
}
