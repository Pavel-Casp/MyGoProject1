package main

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Numbers []float64 `json:"numbers"`
}

type Response struct {
	Result float64 `json:"result"`
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var sum float64
	for _, n := range req.Numbers {
		sum += n
	}

	resp := Response{Result: sum}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/sum", sumHandler)
	http.ListenAndServe(":8080", nil)
}
