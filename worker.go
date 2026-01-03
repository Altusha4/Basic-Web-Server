package main

import (
	"context"
	"log"
	"time"
)

func StartBackgroundWorker(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Timetable server is running")

		case <-ctx.Done():
			log.Println("Background worker stopped")
			return
		}
	}
}
