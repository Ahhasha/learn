package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	tom := new(person)

	fmt.Println(*tom)

	tom.name = "Tom"
	tom.age = 41
	fmt.Println(*tom)

	bob := new(struct {
		name, company string
		age           int
	})

	fmt.Println(*bob)

	bob.name = "Bobby"
	bob.company = "Avito"
	bob.age = 22

	fmt.Println(*bob)
}
