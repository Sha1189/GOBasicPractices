package main

import "fmt"

func main() {
	var a [10]int
	fmt.Println("Array: ", a)

	a[1] = 5
	fmt.Println("Array: ", a)
	fmt.Println(a[1])

	declaredArrayNoLength := [...]int{0, 1, 2, 3}
	declaredArrayWithLength := [10]int{0, 1, 2, 3} // ... can be replace with numerical value
	fmt.Println(declaredArrayNoLength)
	fmt.Println(declaredArrayWithLength)

	fmt.Println(len(declaredArrayNoLength))
	fmt.Println(len(declaredArrayWithLength))
}
