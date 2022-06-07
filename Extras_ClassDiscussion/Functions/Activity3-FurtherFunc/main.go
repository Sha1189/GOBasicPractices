package main

func main() {
	var (
		minecraft       = game{title: "Minecraft", price: 5}
		worldOfWarcraft = game{title: "World of Warcraft", price: 19}
		eliteDangerous  = game{title: "Elite Dangerous", price: 54}
	)

	var items []game
	items = append(items, minecraft, worldOfWarcraft, eliteDangerous)

	my := list(items)
	my.print()
}
