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

func (s *TimetableStorage) GetAll() map[string]TimetableEntry {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(map[string]TimetableEntry)
	for k, v := range s.data {
		result[k] = v
	}
	return result
}

func (s *TimetableStorage) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[id]; ok {
		delete(s.data, id)
		return true
	}
	return false
}
