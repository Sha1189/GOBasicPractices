package main

import "fmt"

const (
	Weight = 44
	Height = 165
	Age    = 22
)

const (
	count0 = iota
	count1
	count2
	count3
)

func main() {
	fmt.Println(Weight)
	fmt.Println(Height)
	fmt.Println(Age)
	fmt.Println(count0)
	fmt.Println(count1)
	fmt.Println(count2)
	fmt.Println(count3)

}
