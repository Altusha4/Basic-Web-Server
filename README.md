# Concurrent Web Server — Go (net/http)

A concurrent RESTful web server built using Go’s `net/http` package.  
The application demonstrates safe concurrent request handling, shared state protection with mutexes, background workers, and graceful shutdown using context.

Developed as part of **Advanced Programming 1 — Assignment 2**.

---

## Technologies Used

- ![](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
- ![](https://img.shields.io/badge/net/http-000000?style=for-the-badge&logo=go&logoColor=white)
- ![](https://img.shields.io/badge/Goroutines-00ADD8?style=for-the-badge&logo=go&logoColor=white)
- ![](https://img.shields.io/badge/Mutex-4B5563?style=for-the-badge)
- ![](https://img.shields.io/badge/Context-2563EB?style=for-the-badge)
- ![](https://img.shields.io/badge/JSON-F59E0B?style=for-the-badge)


---

## Features

- RESTful API built with Go `net/http`
- Concurrent request handling (goroutines)
- Thread-safe in-memory storage using mutex
- Background worker running periodically
- Graceful server shutdown (Ctrl+C)
- JSON-based request and response format
- Simple and clean project structure
- Real-world data model (`TimetableEntry`)

---

## API Endpoints

| Method | Endpoint | Description |
|------|--------|------------|
| POST | `/data` | Add a timetable entry (JSON) |
| GET | `/data` | Get all stored entries |
| DELETE | `/data/{key}` | Delete entry by key (ID) |
| GET | `/stats` | Get server statistics |

---

## Data Model

The server stores structured JSON data using the following model:

```json
{
  "id": "SE-2416",
  "subject": "Advanced Programming",
  "day": "Wednesday",
  "time": "14:00-15:50",
  "room": "С1.1.239",
  "teacher": "Nurlybek T."
}
```

## Project Structure

.
├── main.go        # Server startup & graceful shutdown  
├── model.go       # TimetableEntry data model  
├── storage.go     # In-memory storage with mutex protection  
├── service.go     # Business logic layer  
├── handler.go     # HTTP handlers & statistics  
├── worker.go      # Background worker (periodic logging)  
├── go.mod         # Go module definition  
└── README.md  

---

## Concurrency & Thread Safety

- net/http automatically handles each request in a separate goroutine  
- Shared resources are protected using:  
  - sync.RWMutex for the in-memory map  
  - sync.Mutex for request counters  
- Prevents race conditions during concurrent access  

---

## Background Worker

- Runs in a separate goroutine  
- Logs server status every 5 seconds  
- Uses time.Ticker and select statement  
- Stops cleanly when shutdown signal is received  


---

## Graceful Shutdown

- Implemented using context.Context  
- OS signals (Ctrl+C, SIGTERM) are captured  
- Server shuts down without interrupting active requests  
- Background worker stops gracefully  

This follows industry-standard practices.

---

## Server Statistics (/stats)

Minimum required:
```json
{
  "total_requests": 10
}
```

# How to Run the Project

## Install Go
Make sure Go is installed:
```
go version
```
Run the server
```
go run .
```
Server address
```
http://localhost:8080
```
Example Usage (curl)
Add data
```
curl -X POST http://localhost:8080/data \
  -H "Content-Type: application/json" \
  -d '{"id":"SE-2416","subject":"ADP1","day":"Wednesday","time":"14:00","room":"C1.1.239","teacher":"Nurlybek"}'
```
Get all data
```
curl http://localhost:8080/data
```
Delete data
```
curl -X DELETE http://localhost:8080/data/CS101
```
Get statistics
```
curl http://localhost:8080/stats
```
## How It Works

### Request Flow
- Client sends an HTTP request  
- `net/http` spawns a new goroutine for the request  
- Handler processes the request logic  
- Mutex protects shared state (map and counters)  
- Response is returned as JSON  

### Background Process
- Background worker runs in a separate goroutine  
- Logs server status every 5 seconds  
- Uses `select-case` for multiplexing  
- Stops cleanly on context cancellation  

---

## Assignment Compliance

✔ net/http web server  
✔ Concurrent request handling  
✔ Mutex-protected shared state  
✔ Background worker (goroutine)  
✔ select-case usage  
✔ Context-based graceful shutdown  
✔ REST API endpoints  
✔ `/stats` endpoint implemented  

Fully complies with Assignment 2 requirements.

---

## Author

Altynay Yertay  
Software Engineering Student  
Astana IT University
