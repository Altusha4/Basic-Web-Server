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

func (h *TimetableHandler) GetTimetable(w http.ResponseWriter, r *http.Request) {
	data := h.service.GetAllEntries()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *TimetableHandler) DeleteTimetable(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	if !h.service.DeleteEntry(id) {
		http.Error(w, "entry not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TimetableHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	requests, entries := h.service.Stats()

	response := map[string]int{
		"total_requests": requests,
		"total_entries":  entries,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
