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

func (s *TimetableService) GetAllEntries() map[string]TimetableEntry {
	s.totalRequests++
	return s.storage.GetAll()
}

func (s *TimetableService) DeleteEntry(id string) bool {
	s.totalRequests++
	return s.storage.Delete(id)
}

func (s *TimetableService) Stats() (int, int) {
	return s.totalRequests, s.storage.Count()
}
