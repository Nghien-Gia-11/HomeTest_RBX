package domain

import "time"

type Log struct {
	Timestamp  time.Time
	PageId     int
	CustomerId int
}
