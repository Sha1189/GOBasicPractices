package main

import "fmt"

type computerAccessories struct {
	item  string
	price float32
}

func (c *computerAccessories) print() {
	fmt.Println(c.item, c.price)
}
