package main

import "fmt"

func sumMap(m map[string]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}
func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("Сумма:", sumMap(m))
	m1 := map[string]int{"fifa": 4, "five": 5, "pepero": 6}
	fmt.Println("Сумма:", sumMap(m1))
}
