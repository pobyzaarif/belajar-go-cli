package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

type Item struct {
	Item  string
	Price float64
}

var Items = []Item{
	{Item: "Big Mac", Price: 5.99},
	{Item: "Quarter Pounder with Cheese", Price: 6.49},
	{Item: "McChicken", Price: 2.99},
	{Item: "Medium Fries", Price: 3.49},
	{Item: "Medium Soft Drink", Price: 2.19},
	{Item: "Chicken McNuggets (10 pc)", Price: 5.29},
	{Item: "Filet-O-Fish", Price: 4.99},
	{Item: "Egg McMuffin", Price: 4.49},
	{Item: "Apple Pie", Price: 1.89},
	{Item: "McFlurry", Price: 3.79},
}

var NameItems = func() (listNameItem []string) {
	for _, v := range Items {
		listNameItem = append(listNameItem, v.Item)
	}
	return listNameItem
}

var mapItemPrice = func() map[string]float64 {
	itemPrice := map[string]float64{}
	for _, v := range Items {
		itemPrice[v.Item] = v.Price
	}
	return itemPrice
}

type Order struct {
	Price    float64
	Qty      float64
	Subtotal float64
}

func main() {
	prompt := promptui.Select{
		Label: "Options",
		Items: []string{"Menu", "New Order", "Exit"},
	}

	for {
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		switch result {
		case "Menu":
			showMenu()
		case "New Order":
			order()
		default:
			fmt.Println("Thank you")
			return
		}
	}
}

func showMenu() {
	data := [][]string{
		{"Item", "Price"},
	}

	for _, v := range Items {
		addData := []string{v.Item, fmt.Sprintf("$%v", v.Price)}
		data = append(data, addData)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	_ = table.Bulk(data[1:])
	_ = table.Render()
}

func order() {
	// create new order
	newOrder := make(map[string]Order)

	fmt.Println("⚠️ Take 10% off your purchase - just spend $10 or more!")

initmenu:
	lastPrompt := promptui.Select{
		Label: "Next action",
		Items: []string{"Add Item to Order", "Check My Order", "Finish"},
	}

	_, resultLast, err := lastPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch resultLast {
	case "Add Item to Order":
		goto chooseItem
	case "Check My Order":
		printOrder(newOrder)
		goto initmenu
	default:
		printOrder(newOrder)
		return
	}

chooseItem:
	itemPrompt := promptui.Select{
		Label: "Choose Items",
		Items: NameItems(),
	}

	_, resultItem, err := itemPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	qtyPrompt := promptui.Prompt{
		Label:    "Qty",
		Validate: validate,
	}

	resultQty, err := qtyPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	qtyFloat, _ := strconv.ParseFloat(resultQty, 64)
	subTotal := mapItemPrice()[resultItem] * qtyFloat

	if _, ok := newOrder[resultItem]; !ok && qtyFloat == 0 {
		goto initmenu
	}

	newOrder[resultItem] = Order{
		Price:    mapItemPrice()[resultItem],
		Qty:      newOrder[resultItem].Qty + qtyFloat,
		Subtotal: newOrder[resultItem].Subtotal + subTotal,
	}
	printOrder(newOrder)
	goto initmenu
}

func printOrder(newOrder map[string]Order) {
	data := [][]string{
		{"Item", "Price", "Qty", "Subtotal"},
	}
	total := 0.00
	for k, v := range newOrder {
		addData := []string{k, fmt.Sprintf("%v", v.Price), fmt.Sprintf("%v", v.Qty), fmt.Sprintf("$%v", v.Subtotal)}
		data = append(data, addData)
		total = total + v.Subtotal
	}

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithConfig(tablewriter.Config{
			Row: tw.CellConfig{
				Formatting:   tw.CellFormatting{AutoWrap: tw.WrapNormal}, // Wrap long content
				Alignment:    tw.CellAlignment{Global: tw.AlignLeft},     // Left-align rows
				ColMaxWidths: tw.CellWidth{Global: 25},
			},
			Footer: tw.CellConfig{
				Alignment: tw.CellAlignment{Global: tw.AlignLeft},
			},
		}),
	)
	table.Header(data[0])
	_ = table.Bulk(data[1:])

	if total > 10.00 {
		discount := 0.1
		discountPrice := total * discount
		total = total - discountPrice

		table.Footer([]string{"", "", "DISCOUNT\nTOTAL", fmt.Sprintf("$%.2f\n$%.2f", discountPrice, total)})
	} else {
		table.Footer([]string{"", "", "TOTAL", fmt.Sprintf("%.2f", total)})
	}
	table.Configure(func(config *tablewriter.Config) {
		config.Row.Alignment.Global = tw.AlignLeft
	})
	_ = table.Render()
}
