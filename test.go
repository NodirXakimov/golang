package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println(name, price)
	case "s":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println(tip)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func main() {
	bill := newBill("Nodir's bill")
	bill.addItem("Soup", 2.55)
	bill.addItem("Hamburger", 3.99)
	bill.addItem("Coke", 1.99)
	bill.updateTip(10)
	fmt.Println(bill.format())
	myBill := createBill()
	promptOptions(myBill)
}
