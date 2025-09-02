package main

import "fmt"

func action(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

func main() {

	action(10, 25, func(x int, y int) int { return x + y })
	action(5, 7, func(x int, y int) int { return x * y })
}
