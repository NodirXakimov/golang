package main

import "fmt"

func main() {
	var str string
	fmt.Print("Enter a word to check palindrome: ")
	fmt.Scanln(&str)
	palindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			palindrome = false
			break
		}
	}
	if palindrome {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}
}
