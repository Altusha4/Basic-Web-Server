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

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	go StartBackgroundWorker(ctx)

	go func() {
		log.Println("Server started on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutdown signal received")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("shutdown error: %v", err)
	}

	log.Println("Server stopped")
}
