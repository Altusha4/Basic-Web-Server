package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type TimetableHandler struct {
	service *TimetableService

	mu       sync.Mutex
	requests int
}

func NewTimetableHandler(service *TimetableService) *TimetableHandler {
	return &TimetableHandler{
		service: service,
	}
}

func (h *TimetableHandler) CreateTimetable(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	h.requests++
	h.mu.Unlock()

	var entry TimetableEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if entry.ID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	h.service.AddEntry(entry)
	w.WriteHeader(http.StatusCreated)
}

func (h *TimetableHandler) GetTimetable(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	h.requests++
	h.mu.Unlock()

	data := h.service.GetAllEntries()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *TimetableHandler) DeleteTimetable(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	h.requests++
	h.mu.Unlock()

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
	h.mu.Lock()
	h.requests++
	req := h.requests
	h.mu.Unlock()

	response := map[string]int{
		"total_requests": req,
		"total_entries":  h.service.Count(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
