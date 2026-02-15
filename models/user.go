package models

import "fmt"

type User struct {
	ID       int
	Name     string
	username string
	password string
}

func (u User) Save() {
	// later: add to db or file
	fmt.Printf("User %s has been saved.", u.Name)
}
