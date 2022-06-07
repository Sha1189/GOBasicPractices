package main

import (
	"fmt"
	"strings"
)

//MyStr is a type of string
type MyStr string

//Uppercase is a method in MyStr
func (s MyStr) Uppercase() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Println(MyStr("Go School").Uppercase())
}
