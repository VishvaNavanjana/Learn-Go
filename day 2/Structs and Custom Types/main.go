package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(promt string, r *bufio.Reader) (string,error) {
	fmt.Print(promt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
} 

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ",reader)

	b := newBill(name)
	fmt.Println("Created the bille - ",b.name)

	return b
}

func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save item, t - add tip): ",reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price,64)

		if err != nil {
			fmt.Println("the price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)

		fmt.Println("Item added - ",name,price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount ($) ", reader)

		t, err := strconv.ParseFloat(tip,64)

		if err != nil {
			fmt.Println("the tip must be a number")
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Println("tip added -",tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("you saved the file - ", b.name)
	default:
		fmt.Println("wrong input")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}