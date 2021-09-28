package main

import (
	"fmt"
)

func main() {
	var slice = []string{"Nodir", "Shokir", "Dilshod", "Husan", "Izzat", "Shohruh"}
	// slice = append(slice, "Sirojiddin")
	// fmt.Println(slice, len(slice))
	// rangeOne := slice[1:3]
	// rangeTwo := slice[3:]
	// fmt.Println(rangeOne, len(rangeOne), len(slice))
	// fmt.Println(rangeTwo, len(rangeTwo), len(slice))

	// greeting := "Hello there! How are you doing?"
	// fmt.Println(strings.Contains(greeting, "ello"))
	// fmt.Println(strings.Index(greeting, "o"))
	// fmt.Println(len(strings.Split(greeting, "")))

	// sort.Strings(slice)
	// fmt.Println(sort.SearchStrings(slice, "Nodir"))

	for index, value := range slice {
		fmt.Printf("The value of the %v element is %v \n", index, value)
	}
}
