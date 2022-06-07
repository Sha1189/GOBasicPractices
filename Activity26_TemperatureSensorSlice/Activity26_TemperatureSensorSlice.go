package main

import "fmt"

func main() {
	room1 := []float64{20, 21, 23, 25, 22}
	room2 := []float64{27, 23, 25, 20, 30, 24}
	room3 := []float64{22, 23, 24, 22}
	sliceInSlice := [][]float64{}
	fmt.Println(len(sliceInSlice))
	fmt.Println(cap(sliceInSlice))
	sliceInSlice = append(sliceInSlice, room1)
	sliceInSlice = append(sliceInSlice, room2)
	sliceInSlice = append(sliceInSlice, room3)
	fmt.Println(sliceInSlice)

	for i := 0; i < len(sliceInSlice); i++ {
		total := 0.0
		for j := 0; j < len(sliceInSlice[i]); j++ {
			total += sliceInSlice[i][j]
		}
		average := total / float64(len(sliceInSlice[i]))
		fmt.Println(average)
	}
}
