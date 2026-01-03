package main

import "sync"

type DataStorage struct {
	mu   sync.RWMutex
	data map[string]TimetableEntry
}

func NewDataStorage() *DataStorage {
	return &DataStorage{
		data: make(map[string]TimetableEntry),
	}
}

func (s *DataStorage) Set(key string, value TimetableEntry) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *DataStorage) GetAll() map[string]TimetableEntry {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(map[string]TimetableEntry, len(s.data))
	for k, v := range s.data {
		result[k] = v
	}
	return result
}

func (s *DataStorage) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[key]; ok {
		delete(s.data, key)
		return true
	}
	return false
}

func (s *DataStorage) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}
