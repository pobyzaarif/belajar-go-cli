package handler

import (
	"context"

	"github.com/pobyzaarif/belajar-go-cli/entity"
)

func (h Handler) GetRevenueReport() (reports []entity.RevenueReport, err error) {
	rows, err := h.DB.QueryContext(context.Background(),
		`SELECT
		f.name AS field_name,
		COUNT(b.id) AS total_bookings,
		COALESCE(SUM(b.total_amount), 0) AS total_revenue
	FROM fields f
	LEFT JOIN bookings b ON b.field_id = f.id
	GROUP BY f.id, f.name
	ORDER BY total_revenue DESC;`)
	if err != nil {
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		var report entity.RevenueReport

		err = rows.Scan(&report.FieldName, &report.TotalBooking, &report.TotalRevenue)
		if err != nil {
			return reports, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func (h Handler) GetCustomerWithoutPayment() (custs []entity.CustomerWithoutPayment, err error) {
	rows, err := h.DB.QueryContext(context.Background(), `
	SELECT
		CONCAT(c.first_name, ' ', c.last_name) AS customer_name,
		b.id AS booking_id,
		b.booking_date
	FROM bookings b
	JOIN customers c ON c.id = b.customer_id
	LEFT JOIN payments p ON p.booking_id = b.id
	WHERE p.id IS NULL
	ORDER BY b.booking_date ASC;

	`)
	if err != nil {
		return custs, err
	}
	defer rows.Close()

	for rows.Next() {
		var cust entity.CustomerWithoutPayment

		err = rows.Scan(&cust.Name, &cust.BookingID, &cust.BookingDate)
		if err != nil {
			return custs, err
		}
		custs = append(custs, cust)
	}

	return custs, err
}
