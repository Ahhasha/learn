package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	tom := person{name: "Tom", age: 22}
	var p_tom *person = &tom

	fmt.Println(p_tom.name)
	fmt.Println(p_tom.age)

	p_tom.age = 23
	fmt.Println(tom.age)
}
