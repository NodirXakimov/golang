package main

import (
	"fmt"
	"os"
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

// save the bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("saved_bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}
