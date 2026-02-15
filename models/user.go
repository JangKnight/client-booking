package models

import (
	"example.com/client-booking/db"
	"example.com/client-booking/utils"
)

type User struct {
	ID       int
	Name     string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	// later: add to db or file
	query := "INSERT INTO user(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = int(userId)

	return err
}
