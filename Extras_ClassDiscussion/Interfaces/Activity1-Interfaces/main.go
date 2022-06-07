package main

func main() {
	var (
		minecraft        = game{title: "Minecraft", price: 5}
		worldOfWarcraft  = game{title: "World of warcraft", price: 19}
		eliteDangerous   = game{title: "Elite Dangerous", price: 54}
		candleInTheTomb  = book{title: "Candle in the Tomb", price: 20}
		barneyAndFriends = book{title: "Barney and Friends", price: 10}
		razerBTEarpiece  = computerAccessories{item: "Razer BT earpiece", price: 159}
		razerKeyboard    = computerAccessories{item: "Razer Keyboard", price: 110}
		logitechMouse    = computerAccessories{item: "Logitech Mouse", price: 80}
	)

	var store list
	store = append(store, minecraft, worldOfWarcraft, eliteDangerous, candleInTheTomb, barneyAndFriends, razerBTEarpiece, razerKeyboard, logitechMouse)
	store.print()
}
