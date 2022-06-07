package main

import "fmt"

func main() {
	slice := make([]string, 10)
	slice[0] = "|1st position|"
	slice[1] = "|2nd position|"
	slice[2] = "|3rd position|"
	slice[3] = "|4rd position|"
	slice[9] = "|10th position|"

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	slice = append(slice, "|11th position|")
	slice = append(slice, "|12th position|")
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
