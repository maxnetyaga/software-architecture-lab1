package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := TimeResponse{Time: time.Now().Format(time.RFC3339)}

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	if err := enc.Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)

	fmt.Println("Starting server on :8795")
	log.Fatal(http.ListenAndServe(":8795", nil))
}