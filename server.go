package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type server struct {
	mux *http.ServeMux
	dataProvider busDataProvider
}

func (s *server) handleGetBusData(w http.ResponseWriter, r *http.Request) {
	if s.dataProvider == nil {
		log.Fatalln("no bus data provider registered")
	}

	w.Header().Set("Content-Type", "application/json")
	busData, err := s.dataProvider.fetch()
	if err != nil {
		resData := map[string]any{
			"code": http.StatusInternalServerError,
			"description": "Could not fetch bus data. This incident has been logged.",
		}
		log.Printf("could not fetch bus data: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resData)
		return
	}
	
	resData := map[string]any{
		"buses": busData,
		"expiry": s.dataProvider.getExpiry(),
	}
	json.NewEncoder(w).Encode(resData)
}

func (s *server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"pong"}`))
}

func (s *server) serve(url string) error {
	log.Printf("starting server at url \"%s\"\n", url)
	return http.ListenAndServe(url, s.mux)
}

func newServer(provider busDataProvider) *server {
	srv := server{}

	mux := http.NewServeMux()
	srv.mux = mux
	srv.dataProvider = provider

	mux.HandleFunc("GET /api/bus/positions", srv.handleGetBusData)
	mux.HandleFunc("GET /api/ping", srv.handlePing)
	
	return &srv
}
