package main

import "fmt"

func main() {

	sliceNew := [...]int{4, 5, 6, 7, 8, 9}
	fmt.Printf("Position 0 and 5 of the new slice is %v and %v\n", sliceNew[0], sliceNew[5])

	for i := 0; i <= len(sliceNew)-1; i++ {
		fmt.Printf("Using a normal for loop at index %v is %v\n", i, sliceNew[i])
	}

	for i, v := range sliceNew {
		fmt.Printf("Using a for range loop at index %v is %v\n", i, v)
	}

	fmt.Printf("Slice of size 3 starting at position 1 is %v\n", sliceNew[1:4])

	welcomeSlice := []string{"Hi", "Welcome to Go School"}
	cpy := make([]string, 3)
	copy(cpy, welcomeSlice)
	fmt.Println("Copy: ", cpy)

	fmt.Println("Length before append: ", len(cpy))
	fmt.Println("Capacity before append: ", cap(cpy))
	cpy = append(cpy, "This is the first append")
	cpy = append(cpy, "This is the second append")
	fmt.Println("Length after 2 append: ", len(cpy))
	fmt.Println("Capacity after 2 append: ", cap(cpy))

}
