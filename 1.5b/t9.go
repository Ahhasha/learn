package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}
func appendNode(head *Node, value int) *Node {
	newNode := &Node{Value: value}

	if head == nil {
		return newNode
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	return head
}

func main() {
	var head *Node

	head = appendNode(head, 1)
	head = appendNode(head, 2)
	head = appendNode(head, 3)

	printList(head)
}
