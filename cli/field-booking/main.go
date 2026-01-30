package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/joho/godotenv/autoload"
	"github.com/manifoldco/promptui"
	"github.com/pobyzaarif/belajar-go-cli/config"
	"github.com/pobyzaarif/belajar-go-cli/entity"
	"github.com/pobyzaarif/belajar-go-cli/handler"
)

var (
	options = []string{"Generate Revenue Report", "List Customers Without Payment", "Exit"}
)

func main() {
	spew.Dump()
	mysqlDSN := os.Getenv("MYSQL_DSN")
	db := config.InitDatabase(mysqlDSN)
	defer db.Close()

	h := handler.NewHandler(db)

	fmt.Println("==================================================")
	fmt.Println("Ronaldo's Mini Soccer Field Booking Admin CLI")
	fmt.Println("==================================================")

	prompt := promptui.Select{
		Label: "Please select an options",
		Items: options,
	}

	for {
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		switch result {
		case options[0]:
			reports, err := h.GetRevenueReport()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			entity.PrintRevenueReportTableWriter(reports)
		case options[1]:
			custs, err := h.GetCustomerWithoutPayment()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			entity.PrintCustomerWithoutPaymentTableWriter(custs)
		default:
			fmt.Println("Thank you")
			return
		}
	}
}
