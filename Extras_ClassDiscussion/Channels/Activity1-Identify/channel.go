package main

import (
	"fmt"
)

func producer(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}
func main() {
	ch := make(chan int)
	go producer(ch) // trigger a concurrent routine to happen and calls “producer” to start
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}
