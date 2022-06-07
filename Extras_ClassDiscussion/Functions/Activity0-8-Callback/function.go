package main

import "fmt"

func main() {
	printSlice([]int{1, 2, 3, 4, 5}, func(n int) { fmt.Println(n) })

	x := returnFunction([]int{1, 2, 3, 4, 5})
	x()

}

// callback, anonymous function, tied to callback, used in GO Advanced
func printSlice(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func returnFunction(numbers []int) func() {
	return func() {
		for i, v := range numbers {
			fmt.Println(i, v)

		}

	}

}
