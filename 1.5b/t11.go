package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	b := Book{Title: "go", Author: "Alice"}
	jsonData, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}
