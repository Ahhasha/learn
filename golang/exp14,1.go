package main

import "fmt"

type person struct {
	string
	company string
	int
}

func main() {

	tom := person{"Tom", "Google", 41}
	fmt.Println(tom)
}
