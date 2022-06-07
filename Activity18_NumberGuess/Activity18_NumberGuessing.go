package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	hiddenNumber := int64(rand.Intn(100))

	i := 0
	for i <= 5 {
		var number string
		fmt.Println("Guess the number (0 to 100): ")
		fmt.Scanln(&number)

		numberValue, _ := strconv.ParseInt(number, 10, 0)

		if numberValue == hiddenNumber {
			fmt.Println("Bingo! You got it correct!")
			break

			//		} else if numberValue > hiddenNumber {
			//				fmt.Println("Too high!")
		} else if numberValue == 101 {
			fmt.Println("You have given up. The number is ", hiddenNumber)
			break
		} else if numberValue > hiddenNumber {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Too low!")
		}
		if i == 5 {
			fmt.Println("Game over. The correct answer is ", hiddenNumber)
			break
		}
		i++

	}

}
