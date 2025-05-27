package main

func main() {
	s := newServer(&busSheetDataProvider{})
	s.serve(":8080")
}
