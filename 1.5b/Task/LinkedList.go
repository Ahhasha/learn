package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value, Next: l.Head}
	l.Head = newNode
}

func (l *LinkedList) Delete(value int) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}
	current := l.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}
	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func main() {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Prepend(0)
	list.Print()
	fmt.Println("Удаляем 1")
	list.Delete(1)
	list.Print()
}
