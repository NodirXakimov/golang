package main

import "fmt"

func main() {
	bill := newBill("Nodir's bill")
	bill.addItem("Soup", 2.55)
	bill.addItem("Hamburger", 3.99)
	bill.addItem("Coke", 1.99)
	bill.updateTip(10)
	fmt.Println(bill.format())
	myBill := createBill()
	fmt.Println(myBill)
}
