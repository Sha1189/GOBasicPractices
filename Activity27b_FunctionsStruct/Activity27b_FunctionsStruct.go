package main

import "fmt"

type Particulars struct {
	firstName string
	lastName  string
	//_         struct{}
}

func main() {
	customer := Particulars{"John", "doe"}
	fmt.Println(customer)
}
