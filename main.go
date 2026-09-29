package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type info struct {
	Service   string  `json:"service"`
	Version   string  `json:"version"`
	Time      string  `json:"time"`
	UptimeSec float64 `json:"uptime_seconds"`
}

var startTime = time.Now()

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! Try /health or /api/info.\n")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(info{
			Service:   "solar-route",
			Version:   "0.1.0",
			Time:      time.Now().Format(time.RFC3339),
			UptimeSec: time.Since(startTime).Seconds(),
		})
	})

	fmt.Printf("Server listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
