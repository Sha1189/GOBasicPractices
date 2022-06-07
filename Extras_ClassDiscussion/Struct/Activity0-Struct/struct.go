package main

import "fmt"

type myNameCard struct {
	name    string
	myEmail string
}

func main() {
	var myNameCard1 myNameCard
	fmt.Println(myNameCard1)

	myNameCard1 = myNameCard{
		name:    "Ben",
		myEmail: "lkh22@np.edu.sg",
	}

	fmt.Println(myNameCard1)

	myNameCard1.name = "Kheng Hian"
	fmt.Println(myNameCard1)
}
