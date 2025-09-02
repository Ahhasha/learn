package main

import "fmt"

func main() {

	tom := struct {
		name string
		age  int
	}{
		name: "Tom",
		age:  41,
	}

	fmt.Println(tom.name)
	fmt.Println(tom.age)
}
