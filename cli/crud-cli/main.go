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

func main() {
	spew.Dump()
	mysqlDSN := os.Getenv("MYSQL_DSN")
	db := config.InitDatabase(mysqlDSN)
	defer db.Close()

	h := handler.NewHandler(db)

	prompt := promptui.Select{
		Label: "Options",
		Items: []string{"List All User", "Add new user", "Exit"},
	}

	for {
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		switch result {
		case "List All User":
			users, err := h.GetAllUsers()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				entity.PrintUserTableWriter(users)
			}
		case "Add new user":
			// get input email
			// get input name
			// get input age

			// call h.CreateUser(email, name, age)
		default:
			fmt.Println("Thank you")
			return
		}
	}
}
