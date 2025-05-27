package main

func main() {
	s := newServer()
	s.serve(":8080")
}
