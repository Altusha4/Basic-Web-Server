package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	storage := NewTimetableStorage()
	service := NewTimetableService(storage)
	handler := NewTimetableHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /timetable", handler.CreateTimetable)
	mux.HandleFunc("GET /timetable", handler.GetTimetable)
	mux.HandleFunc("DELETE /timetable/{id}", handler.DeleteTimetable)
	mux.HandleFunc("GET /stats", handler.GetStats)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

}
