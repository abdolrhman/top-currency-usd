package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"top-coins-task/src/api/db"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		limit := r.URL.Query()["limit"][0]
		if limit == "" {
			limit = "10"
		}
		pgDb := db.Connect()
		intLimit, _ := strconv.Atoi(limit)
		result := handler(pgDb, intLimit)
		newj, _ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")

		w.Write(newj)
	})
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalf("Server Failed to start: %v", err)
	}
}
