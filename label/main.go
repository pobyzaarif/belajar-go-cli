package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var Menu = map[string]float64{
	"cough syrup":  49750,
	"pain killers": 129500,
	"antibiotics":  225000,
	"antacids":     275500,
	"vitamins":     285000,
}

type OrderItem struct {
	Price float64
	Qty   int
}

func main() {
	spew.Dump("")
	fmt.Println("Welcome to the Pharmacy Bill System!")
	printMenu()
	createOrder()
}

var caser = cases.Title(language.English)

func printMenu() {
	for k := range Menu {
		fmt.Println(caser.String(k))
	}
}

func createOrder() {
	newOrder := make(map[string]OrderItem)

questionItem:
	fmt.Println("What item would you like to order?")
	var inputItem string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		inputItem = strings.TrimSpace(line)
	}

	inputItemLower := strings.ToLower(inputItem)
	if Menu[inputItemLower] <= 0 {
		fmt.Println("Item not available. Please select a valid item from the menu")
		goto questionItem
	}

questionQty:
	fmt.Println("How many units?")
	var inputQty string
	_, _ = fmt.Scanln(&inputQty)
	qty, err := strconv.Atoi(inputQty)
	if err != nil || qty == 0 {
		fmt.Println("Invalid number of units. Please enter a positive number of units in numeric format")
		goto questionQty
	}

	if qty > 100 {
		fmt.Println("Number of units too large. Please enter a reasonable number of units (e.g., less than 100)")
		goto questionQty
	}

	newOrder[inputItemLower] = OrderItem{
		Price: Menu[inputItemLower],
		Qty:   newOrder[inputItemLower].Qty + qty,
	}
	fmt.Printf("Added %d units %s to your order. ", qty, caser.String(inputItem))

	calculateNPrintOrder(newOrder)

fuQuestion:
	fmt.Println("Would you like to order another item would you like to order? (yes/no)")
	var inputNext string
	_, _ = fmt.Scanln(&inputNext)
	inputNextLower := strings.ToLower(inputNext)
	switch inputNextLower {
	case "yes":
		goto questionItem
	case "no":
		summaryNPrintOrder(newOrder)
	default:
		goto fuQuestion
	}

	fmt.Println("Thank you for your order!")
}

func calculateNPrintOrder(orders map[string]OrderItem) {
	total := 0.00

	for _, v := range orders {
		total = total + (float64(v.Qty) * v.Price)
	}

	fmt.Printf("Your total now is %.0f\n", math.Ceil(total))
}

func summaryNPrintOrder(orders map[string]OrderItem) {
	total := 0.00

	for _, v := range orders {
		total = total + (float64(v.Qty) * v.Price)
	}

	tax := math.Ceil(total * 0.0008)
	totalWithDiscount := math.Ceil(total - (total * 0.01))
	if total < 1000000 {
		fmt.Printf("Your total now is %.0f\n", math.Ceil(total))
	} else if total >= 2000000 {
		fmt.Printf("Congratulations! You qualify for a 10%% discount. Your total after discount is %0.f\n", totalWithDiscount)
		fmt.Printf("With tax applied, Your final total is %0.f\n", totalWithDiscount-tax)
		fmt.Println("Additionally, you qualify for a free unit of Antacids. Enjoy your additional item!")
	} else if total >= 1000000 {
		fmt.Printf("Congratulations! You qualify for a 10%% discount. Your total after discount is %0.f\n", totalWithDiscount)
		fmt.Printf("With tax applied, Your final total is %0.f\n", totalWithDiscount-tax)
	}

}
