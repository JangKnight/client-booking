package models

import (
	"example.com/client-booking/db"
)

type User struct {
	ID       int
	Name     string
	email    string `binding:"required"`
	password string `binding:"required"`
}

func (u User) Save() error {
	// later: add to db or file
	query := "INSERT INTO user(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.email, u.password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = int(userId)

	return err
}
