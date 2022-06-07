package main

import "fmt"

type Particulars struct {
	firstName string
	lastName  string
}

type ParticularsEnforce struct {
	firstName string
	lastName  string
	_         int
}

func main() {
	customer := Particulars{"John", "Doe"}

	fmt.Println(customer)
	customer1 := ParticularsEnforce{"John", "Doe"}
	fmt.Println(customer1)
}
