package main

import (
	"log"
	"net/http"
)

type server struct {
	mux *http.ServeMux
	dataProvider *busDataProvider
}

func (s *server) handleGetBusData(w http.ResponseWriter, r *http.Request) {
	if s.dataProvider == nil {
		log.Fatalln("no bus data provider registered")
	}
}

func (s *server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"pong"}`))
}

func (s *server) serve(url string) error {
	log.Printf("starting server at url \"%s\"\n", url)
	return http.ListenAndServe(url, s.mux)
}

func newServer() *server {
	srv := server{}

	mux := http.NewServeMux()
	srv.mux = mux

	mux.HandleFunc("GET /{$}", srv.handleGetBusData)
	mux.HandleFunc("GET /ping", srv.handlePing)
	
	return &srv
}
