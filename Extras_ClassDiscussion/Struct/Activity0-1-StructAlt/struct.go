package main

import "fmt"

type myProfile struct {
	myName string
}

type myNameCard struct {
	myProfile
	myEmail string
}

func main() {
	myNameCard1 := myNameCard{
		myProfile: myProfile{
			"Ben",
		},
		myEmail: "lkh22@np.edu.sg",
	}

	fmt.Println(myNameCard1)
	fmt.Println(myNameCard1.myProfile.myName)
	fmt.Println(myNameCard1.myName)
}
