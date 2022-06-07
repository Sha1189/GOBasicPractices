package main

import "fmt"

func main() {
	switch {
	case 20 > 40:
		fmt.Println("20> 40")
	case 10 > 1:
		fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 1")
	}

	switch {
	case 10 > 11:
		fmt.Println("10> 11")
	case 5 > 1:
		//fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 2")
	}

	switch {
	case 20 > 21:
		fmt.Println("20>21")
	case 10 > 1:
		break
	case 20 > 10:
		fmt.Println("20>10")
	default:
		fmt.Println("None 3")
	}
}
