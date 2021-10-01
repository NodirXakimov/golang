package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Create a bill
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created bill name: ", b.name)
	return b
}

// Get users input
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option: (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added - ", t)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file - ", b.name)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
