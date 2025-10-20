package main

import "time"

type Log struct {
	Timestamp  time.Time
	PageId     int
	CustomerId int
}
