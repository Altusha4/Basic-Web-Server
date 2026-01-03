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
	storage := NewDataStorage()
	service := NewDataService(storage)
	handler := NewDataHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /data", handler.PostData)
	mux.HandleFunc("GET /data", handler.GetData)
	mux.HandleFunc("DELETE /data/{key}", handler.DeleteData)
	mux.HandleFunc("GET /stats", handler.GetStats)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go StartBackgroundWorker(ctx, handler)

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