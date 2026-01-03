package main

type TimetableService struct {
	storage *TimetableStorage
}

func NewTimetableService(storage *TimetableStorage) *TimetableService {
	return &TimetableService{
		storage: storage,
	}
}

func (s *TimetableService) AddEntry(entry TimetableEntry) {
	s.storage.Add(entry)
}

func (s *TimetableService) GetAllEntries() map[string]TimetableEntry {
	return s.storage.GetAll()
}

func (s *TimetableService) DeleteEntry(id string) bool {
	return s.storage.Delete(id)
}

func (s *TimetableService) Count() int {
	return s.storage.Count()
}
