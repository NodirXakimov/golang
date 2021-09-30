package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make a new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// format the bill
func (b bill) format() string {
	fs := "\nBill breakdown: \n"
	var total float64 = 0

	// list the items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "Tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%v\n", "Total:", total+b.tip)
	return fs
}

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
