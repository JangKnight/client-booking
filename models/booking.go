package models

import (
	"fmt"
	"time"
)

type Booking struct {
	ID            int
	CustomerEmail string `binding:"required"`
	CustomerName  string
	Service       string
	Location      string
	DateTime      time.Time
	UserID        int
}

func (b Booking) Save() {
	// later: add to db or file
	fmt.Printf("Booking for %s has been saved.", b.CustomerName)
}

func (b Booking) GetBookingsByUserID(userID int) []Booking {
	// later: retrieve from db or file
	return []Booking{
		{ID: 1, CustomerEmail: "customer1@example.com", CustomerName: "John Doe", Service: "Consultation", Location: "REMOTE", DateTime: time.Now(), UserID: 1},
		{ID: 2, CustomerEmail: "customer2@example.com", CustomerName: "Jane Smith", Service: "Training", Location: "Los Angeles", DateTime: time.Now().AddDate(0, 0, 1), UserID: 1},
	}
}
