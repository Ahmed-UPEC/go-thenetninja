package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)
	return input, err
}

func createBill() bill {

	reader := bufio.NewReader(os.Stdin)
	name, _ := getUserInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getUserInput("Choose options (a - add an item, s - save the bill, t - add tip, f - format): ", reader)

	switch opt {
	case "a":
		name, _ := getUserInput("Item name: ", reader)
		price, _ := getUserInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		} else {
			b.addItem(name, p)
			fmt.Printf("Adding %v to %v bill... \n", name, b.name)
			promptOptions(b)
		}
	case "t":
		tip, _ := getUserInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		} else {
			b.updateTip(t)
			fmt.Printf("Adding %v tip to %v bill \n", t, b.name)
			promptOptions(b)
		}
	case "s":
		b.save()
		fmt.Println("You save the bill -", b.name)
	case "f":
		fmt.Println(b.format())
		promptOptions(b)
	default:
		fmt.Println("That was not an option")
		promptOptions(b)
	}

}

func main() {

	mybill := createBill()
	promptOptions(mybill)

}
