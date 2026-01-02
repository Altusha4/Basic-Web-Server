package main

type TimetableService struct {
	storage *TimetableStorage

	totalRequests int
}

func NewTimetableService(storage *TimetableStorage) *TimetableService {
	return &TimetableService{
		storage: storage,
	}
}

func (s *TimetableService) AddEntry(entry TimetableEntry) {
	s.totalRequests++
	s.storage.Add(entry)
}
