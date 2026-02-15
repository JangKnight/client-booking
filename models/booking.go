package models

import "time"

type Booking struct {
	ID           int
	CustomerName string
	Service      string
	Location     string
	DateTime     time.Time
	UserID       int
}
