package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.WriteFile("file.txt", []byte("Hello, World!"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
