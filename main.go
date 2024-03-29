package main

import (
	"fmt"
	"github.com/nathandennis/tmcheckout/checkout"
)

type Checkout interface {
	Scan(items ...string) []string
	CalculateTotalPrice() int
}

func main() {
	stock := map[string]checkout.Item{
		"A": {SKU: "A", UnitPrice: 50, MultiBuyQuantity: 3, SpecialPrice: 130},
		"B": {SKU: "B", UnitPrice: 30, MultiBuyQuantity: 2, SpecialPrice: 45},
		"C": {SKU: "C", UnitPrice: 20},
		"D": {SKU: "D", UnitPrice: 15},
	}

	var co Checkout = checkout.New(stock)

	// scan a list with some nonexistent SKUs to check scanner logic
	unrecognizedSKUs := co.Scan("A", "B", "B", "A", "A", "B", "C", "D", "Z", "Z", "T", "O", "P")

	// expect total to be 240 in this instance
	total := co.CalculateTotalPrice()

	fmt.Println("total price: ", total)

	if len(unrecognizedSKUs) > 0 {
		fmt.Println("unrecognized SKUs: ", unrecognizedSKUs)
	}
}
