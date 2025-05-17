package domain

import "time"

type Transaction struct {
	Amount int
	Date   time.Time
}
