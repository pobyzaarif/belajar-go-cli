package entity

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type RevenueReport struct {
	FieldName    string
	TotalBooking int
	TotalRevenue float64
}

func PrintRevenueReportTableWriter(reports []RevenueReport) {
	data := [][]string{
		{"Field Name", "Total Booking", "Total Revenue"},
	}

	for _, v := range reports {
		addData := []string{v.FieldName, fmt.Sprintf("%v", v.TotalBooking), fmt.Sprintf("%v", v.TotalRevenue)}
		data = append(data, addData)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	_ = table.Bulk(data[1:])
	_ = table.Render()
}

type CustomerWithoutPayment struct {
	Name        string
	BookingID   int
	BookingDate time.Time
}

func PrintCustomerWithoutPaymentTableWriter(customers []CustomerWithoutPayment) {
	data := [][]string{
		{"Customer Name", "Booking ID", "Booking Date"},
	}

	for _, v := range customers {
		addData := []string{v.Name, fmt.Sprintf("%v", v.BookingID), v.BookingDate.Format("2006-01-02")}
		data = append(data, addData)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	_ = table.Bulk(data[1:])
	_ = table.Render()
}
