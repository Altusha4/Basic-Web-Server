package main

type TimetableHandler struct {
	service *TimetableService
}

func NewTimetableHandler(service *TimetableService) *TimetableHandler {
	return &TimetableHandler{
		service: service,
	}
}
