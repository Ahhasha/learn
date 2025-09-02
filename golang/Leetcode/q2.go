package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers)
}
