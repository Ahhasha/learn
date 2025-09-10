package main

import (
	"fmt"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	formatted := currentTime.Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Текущее времячко: %s", formatted)
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8080", nil)
}
