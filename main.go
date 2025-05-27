package main

func main() {
	s := newServer(&dummyDataProvider{})
	s.serve(":8080")
}
