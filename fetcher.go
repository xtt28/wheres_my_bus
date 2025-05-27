package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

const busSheet = "https://docs.google.com/spreadsheets/d/1S5v7kTbSiqV8GottWVi5tzpqLdTrEgWEY4ND4zvyV3o/edit"

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
