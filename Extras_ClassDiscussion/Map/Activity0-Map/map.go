package main

import "fmt"

func main() {
	goSchRec := make(map[string]int, 3)

	goSchRec = map[string]int{
		"John":  123,
		"Tom":   234,
		"Jerry": 432,
	}

	fmt.Println(goSchRec)

	seanID := goSchRec["Sean"]
	fmt.Println(seanID)

	seanID, ok := goSchRec["Sean"]
	fmt.Println(seanID, ok)

	goSchRec["John"] = 442
	fmt.Println(goSchRec)

	delete(goSchRec, "Tom")
	fmt.Println(goSchRec)
}
