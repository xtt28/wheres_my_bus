package main

import "os"

func main() {
	s := newServer(&busSheetDataProvider{})
	url := os.Getenv("SERVICE_URL")
	if url == "" {
		url = ":8080"
	}
	s.serve(url)
}
