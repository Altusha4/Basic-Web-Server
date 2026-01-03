package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type DataHandler struct {
	service *DataService

	mu       sync.Mutex
	requests int
}

func NewDataHandler(service *DataService) *DataHandler {
	return &DataHandler{service: service}
}

func (h *DataHandler) incRequests() {
	h.mu.Lock()
	h.requests++
	h.mu.Unlock()
}

func (h *DataHandler) getRequests() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.requests
}

func (h *DataHandler) PostData(w http.ResponseWriter, r *http.Request) {
	h.incRequests()

	var entry TimetableEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if entry.ID == "" {
		http.Error(w, "id is required (used as key)", http.StatusBadRequest)
		return
	}

	h.service.SaveEntry(entry)
	w.WriteHeader(http.StatusCreated)
}

func (h *DataHandler) GetData(w http.ResponseWriter, r *http.Request) {
	h.incRequests()

	data := h.service.GetAll()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func (h *DataHandler) DeleteData(w http.ResponseWriter, r *http.Request) {
	h.incRequests()

	key := r.PathValue("key")
	if key == "" {
		http.Error(w, "missing key", http.StatusBadRequest)
		return
	}

	if !h.service.Delete(key) {
		http.Error(w, "key not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *DataHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	h.incRequests()

	resp := map[string]int{
		"total_requests": h.getRequests(),
		"db_size":        h.service.Count(),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *DataHandler) StatsSnapshot() (int, int) {
	return h.getRequests(), h.service.Count()
}