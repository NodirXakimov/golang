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
	hasDublicate := false
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] == arr[j] {
				hasDublicate = true
				break
			}
		}
	}
	if hasDublicate {
		fmt.Println("There are dublicate elements in array")
	} else {
		fmt.Println("There is no any dublicate elements in array")
	}
}
