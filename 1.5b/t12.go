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
	data := `{"Title":"Go","Author":"Alice"}`
	var b Book
	err := json.Unmarshal([]byte(data), &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b.Title, b.Author)
}
