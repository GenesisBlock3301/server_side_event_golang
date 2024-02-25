package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		// Set headers for server-sent events
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Send server-sent events
		for i := 0; i < 5; i++ { // Send 5 events for demonstration
			fmt.Fprintf(w, "data: %s\n\n", time.Now().Format(time.RFC3339))
			w.(http.Flusher).Flush()    // Flush the response writer to ensure data is sent immediately
			time.Sleep(1 * time.Second) // Wait for 1 second before sending the next event
		}
	})

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
