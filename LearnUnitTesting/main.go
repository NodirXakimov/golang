package main

import "fmt"

func calculate(n int) (result int) {
	result = n + 2
	return
}

func main() {
	fmt.Println("Hello testing...")
	fmt.Println(calculate(13))
}
