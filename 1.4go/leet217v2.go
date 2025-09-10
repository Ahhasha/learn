package main

import "fmt"

func containsDuplicate(nums []int) bool {
	x := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := x[num]; ok {
			return true
		}
		x[num] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 4}))
}
