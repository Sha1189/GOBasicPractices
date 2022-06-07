package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	FName        string
	LName        string
	UserID       int
	InvoiceTotal float64
}

func inspect(n interface{}) {
	refType := reflect.TypeOf(n)
	refValue := reflect.ValueOf(n)
	fmt.Println("Num of fields: ", refType.NumField())
	for i := 0; i < refType.NumField(); i++ {
		fmt.Println(refType.Field(i).Name, "value: ", refValue.Field(i), "type: ", refType.Field(i).Type)
	}
}

func main() {

	customer1 :=
		customer{
			FName:        "John",
			LName:        "Wick",
			UserID:       123123123,
			InvoiceTotal: 10000,
		}
	inspect(customer1)
}
