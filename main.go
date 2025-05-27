package main

func main() {
	busSheetFetchData()
	s := newServer(&busSheetDataProvider{})
	s.serve(":8080")
}
