package model

import "time"

type OrderProcess struct {
	ID          uint16
	DriverID    int
	OrderID     int
	Status      string
	ProcessedAt time.Time
}
