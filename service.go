package main

type DataService struct {
	storage *DataStorage
}

func NewDataService(storage *DataStorage) *DataService {
	return &DataService{storage: storage}
}

func (s *DataService) SaveEntry(entry TimetableEntry) {
	s.storage.Set(entry.ID, entry)
}

func (s *DataService) GetAll() map[string]TimetableEntry {
	return s.storage.GetAll()
}

func (s *DataService) Delete(key string) bool {
	return s.storage.Delete(key)
}

func (s *DataService) Count() int {
	return s.storage.Count()
}