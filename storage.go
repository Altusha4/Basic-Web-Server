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

func (s *TimetableStorage) Add(entry TimetableEntry) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[entry.ID] = entry
}
