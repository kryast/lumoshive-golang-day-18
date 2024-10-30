package model

import "time"

type Orders struct {
	ID          uint16
	PickUp      string
	Destination string
	Price       int
	CreatedAt   time.Time
	CustomerID  int
}
