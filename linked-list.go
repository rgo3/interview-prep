package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Root *Node
}

func (l *List) Insert(elem *Node) error {

	if elem == nil {
		return fmt.Errorf("There is no element to insert!")
	}

	if l.Root == nil {
		l.Root = elem
		return nil
	}

	node := l.Root

	for node.Next != nil {
		node = node.Next
	}

	node.Next = elem

	return nil
}

// (5) -> (6) -> (10)
// <Empty>

func (l *List) PrintList() {

	if l.Root == nil {
		fmt.Println("<Empty>")
		return
	}

	node := l.Root

	fmt.Printf("(%d)", node.Data)

	for node.Next != nil {

		if node.Next != nil {
			fmt.Print(" -> ")
		}

		fmt.Printf("(%d)", node.Next.Data)

		node = node.Next

	}
	fmt.Println()
}

func (l *List) Reverse() {
	node := l.Root
	var previous *Node
	var next *Node

	for node != nil {
		next = node.Next
		node.Next = previous

		previous = node
		node = next
	}

	l.Root = previous
}
