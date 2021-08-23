package entity

import "time"

type Transfer struct {
	ID                   int64
	AccountOriginID      int64
	AccountDestinationID int64
	Amount               float64
	CreateAt             time.Time
}
