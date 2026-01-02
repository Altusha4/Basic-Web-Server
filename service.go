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
