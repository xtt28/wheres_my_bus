package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

const busSheet = "https://docs.google.com/spreadsheets/d/1S5v7kTbSiqV8GottWVi5tzpqLdTrEgWEY4ND4zvyV3o/edit"

func isWithinTimePeriod(subject string, first string, last string) bool {
	subTime, _ := time.Parse("15:04", subject)
	firstTime, _ := time.Parse("15:04", first)
	lastTime, _ := time.Parse("15:04", last)

	return firstTime.Before(subTime) && lastTime.After(subTime)
}

// Require compliance with interface.
var _ busDataProvider = &busSheetDataProvider{}

type busSheetDataProvider struct {
	data busLocationSheetDTO
}

func (p *busSheetDataProvider) fetch() (busLocationSheetDTO, error) {
	if time.Now().Before(p.getExpiry()) && p.data != nil {
		return p.data, nil
	}

	data, err := busSheetFetchData()
	p.data = data
	return data, err
}

func (p *busSheetDataProvider) getExpiry() time.Time {
	now := time.Now()
	timeDayFormatted := now.Format("17:06")

	// Reduce cache TTL during dismissal time period
	if isWithinTimePeriod(timeDayFormatted, "12:25", "12:50") || isWithinTimePeriod(timeDayFormatted, "4:05", "4:30") {
		return now.Add(1 * time.Minute)
	}
	
	return now.Add(5 * time.Minute)
}

// https://www.zenrows.com/blog/golang-html-parser#parse-html-with-the-node-parsing-api-recommended
func traverseTableCells(node *html.Node) []string {
	values := []string{}
	if node.Type == html.ElementNode && node.Data == "td" {
		if node.FirstChild != nil && node.FirstChild.Type == html.TextNode {
			values = append(values, node.FirstChild.Data)
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		values = append(values, traverseTableCells(c)...)
	}

	return values
}

func busSheetFetchData() (busLocationSheetDTO, error) {
	res, err := http.Get(busSheet)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	cells := traverseTableCells(doc)

	fmt.Println(cells)
	return nil, nil
}
