package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"priorityqueue/priorityqueue"
	"strconv"
	"sync"
)

type Window struct {
	ID       int
	IsFree   bool
	Assigned *priorityqueue.Item
}

var (
	queue      = priorityqueue.NewPriorityQueue(2)
	windows    = []*Window{}
	windowLock = sync.Mutex{}
)

func main() {
	const port = 8080

	fmt.Printf("Starting server on %v\n", port)

	for i := 0; i < 3; i++ {
		windows = append(windows, &Window{ID: i + 1, IsFree: true})
	}
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/generate_ticket", generateTicket)
	http.HandleFunc("/assign_ticket", assignTicket)
	http.HandleFunc("/release_window", releaseWindow)

	http.ListenAndServe(":8080", nil)
}

func generateTicket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var jsonC struct {
		IsThirdAge bool `json:"is_third_age"`
	}

	if err := json.NewDecoder(r.Body).Decode(&jsonC); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	queue.Push("Cliente", jsonC.IsThirdAge)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ticket_generated"})
}

func assignTicket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	windowLock.Lock()
	defer windowLock.Unlock()

	for _, window := range windows {
		if window.IsFree {
			ticket := queue.Pop()
			if ticket == nil {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{"status": "no_tickets"})
				return
			}
			window.Assigned = ticket
			window.IsFree = false
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":    "ticket_assigned",
				"window_id": window.ID,
				"ticket":    ticket,
			})
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "no_free_windows"})
}

func releaseWindow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid window id", http.StatusBadRequest)
		return
	}

	windowLock.Lock()
	defer windowLock.Unlock()

	for _, window := range windows {
		if window.ID == id {
			window.IsFree = true
			window.Assigned = nil
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"status": "window_released"})
			return
		}
	}

	http.Error(w, "Invalid window id", http.StatusBadRequest)
}
