package models

import "fmt"

type User struct {
	ID       int
	Name     string
	username string
	password string
}

func (u User) Save() {
	fmt.Printf("User %s has been saved.", u.Name)
}
