package main

import "fmt"

func main() {
	var (
		minecraft       = game{title: "minecraft", price: 5}
		worldOfWarcraft = game{title: "World of warcraft", price: 19}
		eliteDangerous  = game{title: "Elite Dangerous", price: 54}
	)

	var items []*game
	items = append(items, &minecraft, &worldOfWarcraft, &eliteDangerous)
	fmt.Println(items)
	fmt.Println(*items[0])
	my := list(items)
	my.print()
}
