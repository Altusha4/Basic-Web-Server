package main

import (
	"encoding/json"
	"net/http"
)

type TimetableHandler struct {
	service *TimetableService
}

func NewTimetableHandler(service *TimetableService) *TimetableHandler {
	return &TimetableHandler{
		service: service,
	}
}

func (h *TimetableHandler) CreateTimetable(w http.ResponseWriter, r *http.Request) {
	var entry TimetableEntry

	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	h.service.AddEntry(entry)
	w.WriteHeader(http.StatusCreated)
}
