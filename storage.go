package main

import "sync"

type TimetableStorage struct {
	mu   sync.RWMutex
	data map[string]TimetableEntry
}

func NewTimetableStorage() *TimetableStorage {
	return &TimetableStorage{
		data: make(map[string]TimetableEntry),
	}
}
