package main

import (
	"errors"

	"github.com/nfx/go-htmltable"
)

const busSheet = "https://docs.google.com/spreadsheets/d/1S5v7kTbSiqV8GottWVi5tzpqLdTrEgWEY4ND4zvyV3o/edit"

func busSheetFetchData() (busLocationSheetDTO, error) {
	page, err := htmltable.NewFromURL(busSheet)
	if err != nil {
		return nil, err
	}

	if page.Len() < 1 {
		return nil, errors.New("table not found on page")
	}
	table := page.Tables[0]
	rows := table.Rows[2:]

	dto := busLocationSheetDTO{}
	
	for _, row := range rows {
		for i := 1; i < 5; i += 2 {
			if row[i] == "" {
				continue
			}
			dto[row[i]] = row[i + 1]
		}
	}
	
	return dto, nil
}
