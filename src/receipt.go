package main

import (
	"fmt"
	"os"
)

type receipt struct {
	clientName string
	items      map[string]float64
	cash       float64
}

// make new receipts
func newReceipt(clientName string) receipt {
	rcp := receipt{
		clientName: clientName,
		items:      map[string]float64{},
		cash:       0,
	}
	return rcp
}

// add item to receipt
func (rcp *receipt) addItem(name string, price float64) {
	rcp.items[name] = price
}

// format the receipt
func (rcp *receipt) format() string {
	fs := "The GO-SHOP - Receipt breakdown:\n"
	fs += "---------------------------------------\n"
	fs += fmt.Sprintf("%20v\n", "ITEMS")

	var subtotal float64 = 0

	// list items
	for k, v := range rcp.items {
		fs += fmt.Sprintf("%-25v $%v\n", k+":", v)
		subtotal += v
	}

	// Total amount after tax
	total := subtotal * 1.13

	fs += "---------------------------------------\n"

	// add subtotal
	fs += fmt.Sprintf("%-25v $%0.2f\n", "Subtotal:", subtotal)

	// add tax
	fs += fmt.Sprintf("%-25v %v\n", "tax:", "13%")

	// add total
	fs += fmt.Sprintf("%-25v $%0.2f\n", "total:", total)
	fs += "---------------------------------------\n"

	// add cash
	fs += fmt.Sprintf("%-25v $%v\n", "cash:", rcp.cash)

	// add change
	fs += fmt.Sprintf("%-25v $%0.2f\n", "change:", rcp.cash-total)
	fs += "---------------------------------------\n"

	return fs
}

// update cash
func (rcp *receipt) updateCash(cash float64) {
	(*rcp).cash = cash
}

// view the receipt
func (rcp *receipt) view() {
	fmt.Println(rcp.format())
}

// save receipt
func (rcp *receipt) save() {
	data := []byte(rcp.format())
	err := os.WriteFile("../receipts/"+rcp.clientName+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("receipt saved to file")
}
