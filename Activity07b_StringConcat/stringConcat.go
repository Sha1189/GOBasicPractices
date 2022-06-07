package main

import "fmt"

func main() {

	string1 := "This is the first string"
	string2 := "This is the second string"

	fmt.Println(string1 + " " + string2)

	fmt.Println(string1)

	string1 += string2

	fmt.Println(string1)
	fmt.Println(string2)

}
