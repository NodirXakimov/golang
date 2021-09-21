package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println("hello world")
	fmt.Println(fibonacci(50))
}
