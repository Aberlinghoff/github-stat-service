package main

import (
    "log"
    "net/http"

    "github.com/Aberlinghoff/github-stat-service/handlers"
)

func main() {
    http.HandleFunc("/health", handlers.HealthHandler)
    http.HandleFunc("/stats", handlers.StatsHandler)

    log.Println("starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}
