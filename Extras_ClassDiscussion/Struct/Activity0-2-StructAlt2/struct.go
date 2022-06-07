package main

import "fmt"

type myProfile struct {
	myName string
}

type myNameCard struct {
	myprofile myProfile
	myEmail   string
}

func main() {

	myNameCard1 := myNameCard{
		myprofile: myProfile{
			"Ben",
		},
		myEmail: "lkh22@np.edu.sg",
	}

	fmt.Println(myNameCard1)
	fmt.Println(myNameCard1.myprofile.myName)

	myNameCard2 := myNameCard{
		myprofile: myProfile{"BEN"},
		myEmail:   "LKH22@NP.EDU.SG",
	}

	fmt.Println(myNameCard2)
}
