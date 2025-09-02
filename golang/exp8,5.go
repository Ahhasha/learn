package main

import "fmt"

func sum(n uint) uint {
	if n == 1 {
		return n
	}
	return n + sum(n-1)
}
func main() {
	fmt.Println(sum(4))
	fmt.Println(sum(5))
	fmt.Println(sum(6))
}
