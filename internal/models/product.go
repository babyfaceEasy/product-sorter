package models

import "time"

type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    time.Time
	SalesCount int
	ViewsCount int
}
