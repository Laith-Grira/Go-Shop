package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createReceipt() receipt {
	reader := bufio.NewReader(os.Stdin)

	clientName, _ := getInput("Create a new bill for: ", reader)

	rcp := newReceipt(clientName)
	fmt.Println("Created the bill -", rcp.clientName)

	return rcp
}

func promptOptions(rcp receipt) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save receipt, p - pay receipt): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number...")
			promptOptions(rcp)
		}
		rcp.addItem(name, p)

		fmt.Println("item added -", name, price)
		promptOptions(rcp)
	case "p":
		cashString, _ := getInput("Enter cash amount ($): ", reader)

		cash, err := strconv.ParseFloat(cashString, 64)
		if err != nil {
			fmt.Println("The cash must be a number...")
			promptOptions(rcp)
		}

		rcp.updateCash(cash)

		fmt.Println("cash has been updated to", cash)
		promptOptions(rcp)
	case "s":
		rcp.save()
		fmt.Println("receipt has been saved as", rcp.clientName)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(rcp)
	}
}

func main() {

	myReceipt := createReceipt()
	promptOptions(myReceipt)

}
