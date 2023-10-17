package domain

import (
	"time"
)

type Orders struct {
	Id_order		int
	Id_user			int
	Id_product		int
	Total			uint
	Note			string
	Status			string
	Order_date		time.Time
}

type Orders_detail struct {
	Id_order		int
	Id_product		int
	Price			uint
	Quantity		int
}