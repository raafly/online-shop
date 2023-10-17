package domain

import "time"

type Payments struct {
	Id_payment		int
	Id_user			int
	Metode			string
	Detail			string
	Date			time.Time
}