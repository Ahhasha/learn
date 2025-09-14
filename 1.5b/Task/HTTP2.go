package main

import (
	"encoding/json"
	"net/http"
)

type Numbers struct {
	Data []int `json:"data"`
}

type ReverseHandler struct{}

func (h ReverseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var nums Numbers
	err := json.NewDecoder(r.Body).Decode(&nums)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, j := 0, len(nums.Data)-1; i < j; i, j = i+1, j-1 {
		nums.Data[i], nums.Data[j] = nums.Data[j], nums.Data[i]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nums)
}

func main() {
	http.ListenAndServe(":8080", ReverseHandler{})
}
