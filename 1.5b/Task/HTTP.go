package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from custom handler!")
}

func main() {
	http.ListenAndServe(":8080", MyHandler{})
}
