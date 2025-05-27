package main

import (
	"log"
	"time"
)

type busDataProvider interface {
	fetch() (busLocationSheetDTO, error)
	getExpiry() time.Time
}

func isWithinTimePeriod(subject string, first string, last string) bool {
	subTime, _ := time.Parse("15:04", subject)
	firstTime, _ := time.Parse("15:04", first)
	lastTime, _ := time.Parse("15:04", last)

	return firstTime.Before(subTime) && lastTime.After(subTime)
}

// Require compliance with interface.
var _ busDataProvider = &busSheetDataProvider{}
var _ busDataProvider = &dummyDataProvider{}

type busSheetDataProvider struct {
	data busLocationSheetDTO
	expiry time.Time
}

func (p *busSheetDataProvider) fetch() (busLocationSheetDTO, error) {
	if !p.getExpiry().IsZero() && time.Now().Before(p.getExpiry()) && p.data != nil {
		return p.data, nil
	}

	log.Println("busSheetDataProvider: fetching data")
	data, err := busSheetFetchData()
	p.data = data
	p.resetExpiry()
	return data, err
}

func (p *busSheetDataProvider) resetExpiry() {
	now := time.Now()
	timeDayFormatted := now.Format("17:06")

	// Reduce cache TTL during dismissal time period
	if isWithinTimePeriod(timeDayFormatted, "12:25", "12:50") || isWithinTimePeriod(timeDayFormatted, "4:05", "4:30") {
		p.expiry = now.Add(1 * time.Minute)
		return
	}
	
	p.expiry = now.Add(5 * time.Minute)
}

func (p *busSheetDataProvider) getExpiry() time.Time {
	return p.expiry
}

type dummyDataProvider struct {}

func (p *dummyDataProvider) fetch() (busLocationSheetDTO, error) {
	return busLocationSheetDTO{
		"Town 1": "A1",
		"Town 2": "A2",
		"Town 3": "A3",
	}, nil
}

func (p *dummyDataProvider) getExpiry() time.Time {
	return time.Now().Add(5 * time.Hour)
}
