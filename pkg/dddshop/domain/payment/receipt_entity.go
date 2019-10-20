package payment

import "time"

type Receipt struct {
	ID        string
	Price     int
	AccountID string
	CreatedAt time.Time
}
