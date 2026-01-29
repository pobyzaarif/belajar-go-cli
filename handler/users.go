package handler

import (
	"context"

	"github.com/pobyzaarif/belajar-go-cli/entity"
)

func (h Handler) CreateUser(email string, name string, age int) (err error) {
	_, err = h.DB.ExecContext(context.Background(), "INSERT INTO users (email,name,age) VALUES (?, ?, ?)", "john@mail.com", "John Doe", 18)
	if err != nil {
		return err
	}

	return err
}

func (h Handler) GetAllUsers() (users []entity.Users, err error) {
	rows, err := h.DB.QueryContext(context.Background(), "SELECT id, email, name, age FROM users LIMIT 10")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.Users

		err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.Age)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, err
}
