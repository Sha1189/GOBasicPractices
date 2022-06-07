package main

import (
	"fmt"
)

func main() {
	var guess int
	//int guess = 0;
	fmt.Println("Enter an integer value: ")
	fmt.Scanln(&guess)
	fmt.Println(&guess)
	fmt.Println(guess)

	guessvalue := 50

	if guess == guessvalue {
		fmt.Println("Well done! Your guess is correct.")
	}
	if guess < guessvalue {
		fmt.Println("Too low, try again next time!")
	}
	if guess > guessvalue {
		fmt.Println("Too high, try again next time!")
	}

}
