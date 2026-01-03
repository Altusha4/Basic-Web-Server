package main

import (
	"context"
	"log"
	"time"
)

func StartBackgroundWorker(ctx context.Context, handler *DataHandler) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			req, size := handler.StatsSnapshot()
			log.Printf("Current requests: %d, Database size: %d", req, size)

		case <-ctx.Done():
			log.Println("Background worker stopped")
			return
		}
	}
}
