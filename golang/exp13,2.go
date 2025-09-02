package main

import "fmt"

func action(n1 int, n2 int, op func(int, int) int) {

	result := op(n1, n2)
	fmt.Println(result)
}

func add(x int, y int) int {

	return x + y
}

func main() {

	var myOperation func(int, int) int = add
	action(10, 25, myOperation)
}
