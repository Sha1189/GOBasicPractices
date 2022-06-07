package main

import "fmt"

func main() {

	myAppointment := struct {
		position string
	}{
		position: "CET Ops",
	} //anonymous struct

	fmt.Println(myAppointment)

}
