package main

import "fmt"

type mile uint

func main() {
	var distance uint = 5
	fmt.Println(distance)
	distance += 5
	fmt.Println(distance)
}
