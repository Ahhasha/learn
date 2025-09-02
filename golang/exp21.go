package main

import "fmt"

func main() {

	people := map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println(people["Tom"])
	fmt.Println(people["Alice"])
	fmt.Println(people["Bob"])
	people["Bob"] = 34
	fmt.Println(people["Bob"])

	if i, ok := people["Sam"]; ok {
		fmt.Println(i)
	}
	delete(people, "Alice")
	fmt.Println(people)
}
