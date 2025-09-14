package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
