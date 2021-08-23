package entity

import "time"

type Transfer struct {
	ID                     int64
	AccountOriginID        int64
	AccountOriginUUID      string
	AccountDestinationID   int64
	AccountDestinationUUID string
	Amount                 float64
	CreateAt               time.Time
}
