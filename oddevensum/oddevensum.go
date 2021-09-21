package main

import "fmt"

func main() {
	var size int
	fmt.Print("Size of array: ")
	fmt.Scan(&size)
	fmt.Println("Enter ", size, " numbers separated by space")
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	var oddSum, evenSum int
	for i := 0; i < size; i++ {
		if arr[i]%2 == 0 {
			evenSum += arr[i]
		} else {
			oddSum += arr[i]
		}
	}
	fmt.Println("Sum of odd numbers  = ", oddSum)
	fmt.Println("Sum of even numbers = ", evenSum)
}
