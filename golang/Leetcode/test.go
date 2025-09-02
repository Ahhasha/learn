package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["two"] = 34
	m["twoo"] = 35
	m["twooo"] = 36
	m["twooooo"] = 37
	value, exists := m["five"]
	if exists {
		fmt.Println("EST PARASHA:", value)
	} else {
		fmt.Println("NIHYA")
	}
}
