package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
