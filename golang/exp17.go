package main

import "fmt"

type person struct {
	name string
	age  int
}

type account struct {
	login    string
	password string
	person
}

func main() {

	tom := account{
		"tom@mail.com",
		"12334455",
		person{"Tom", 41},
	}

	fmt.Println(tom)

	fmt.Println("Name: ", tom.person.name)
}
