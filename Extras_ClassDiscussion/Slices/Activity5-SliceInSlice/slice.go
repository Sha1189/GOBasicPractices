package main

import "fmt"

func main() {
	room1 := []float64{20, 21, 23, 25, 22}
	room2 := []float64{27, 23, 25, 20, 30, 24}
	room3 := []float64{22, 23, 24, 22}
	/*
		sliceInSlice[0] is room1
		sliceInSlice is room1 , room2 , room3
		len(sliceInSlice[0]) is the length of room1 - number of elements in room 1
	*/
	sliceInSlice := [][]float64{}
	sliceInSlice = append(sliceInSlice, room1)
	sliceInSlice = append(sliceInSlice, room2)
	sliceInSlice = append(sliceInSlice, room3)
	fmt.Printf("Slice in a slice: %v\n", sliceInSlice)

	for i := 0; i < len(sliceInSlice); i++ {
		total := 0.0
		for j := 0; j < len(sliceInSlice[i]); j++ {
			total += sliceInSlice[i][j]
		}
		average := total / float64(len(sliceInSlice[i]))
		fmt.Printf("Average of room %v is %.2f\n", i+1, average)
	}

	for i := range sliceInSlice {
		total := 0.0
		for j := range sliceInSlice[i] {
			total += sliceInSlice[i][j]
		}
		average := total / float64(len(sliceInSlice[i]))
		fmt.Printf("Average of room %v is %.2f\n", i+1, average)
	}

}
