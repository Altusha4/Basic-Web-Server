package main

import "sync"

type TimetableStorage struct {
	mu   sync.RWMutex
	data map[string]TimetableEntry
}
