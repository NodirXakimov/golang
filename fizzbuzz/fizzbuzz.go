package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&n)
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if n%3 == 0 {
		fmt.Println("Fizz")
	} else if n%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("Not dividable by 3 nor 5")
	}
}
