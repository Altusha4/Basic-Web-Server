package main

type TimetableEntry struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Day     string `json:"day"`
	Time    string `json:"time"`
	Room    string `json:"room"`
	Teacher string `json:"teacher"`
}