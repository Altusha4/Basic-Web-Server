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
}
