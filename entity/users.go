package entity

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type Users struct {
	ID    int64
	Email string
	Name  string
	Age   int
}

func PrintUserTableWriter(users []Users) {
	data := [][]string{
		{"ID", "Email", "Name", "Age"},
	}

	for _, v := range users {
		addData := []string{fmt.Sprintf("%v", v.ID), v.Email, v.Name, fmt.Sprintf("%v", v.Age)}
		data = append(data, addData)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	_ = table.Bulk(data[1:])
	_ = table.Render()
}
