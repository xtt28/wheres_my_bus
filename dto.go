package main

import "time"

type busLocationSheetDTO map[string]string

type busDataProvider interface {
	fetch() (busLocationSheetDTO, error)
	getExpiry() time.Time
}
