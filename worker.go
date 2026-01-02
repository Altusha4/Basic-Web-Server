package main

import (
	"context"
	"log"
	"time"
)

func StartBackgroundWorker(
	ctx context.Context,
	service *TimetableService,
) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			requests, entries := service.Stats()
			log.Printf(
				"Timetable status: requests=%d, entries=%d",
				requests,
				entries,
			)

		case <-ctx.Done():
			log.Println("Background worker stopped")
			return
		}
	}
}
