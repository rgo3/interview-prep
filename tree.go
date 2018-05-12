package main

import (
	"fmt"
)

type Leave struct {
	Data  int
	Left  *Leave
	Right *Leave
}

type Tree struct {
	Root *Leave
}

func (t *Tree) Insert(l *Leave) error {
	if l == nil {
		return fmt.Errorf("no element to insert")
	}

	if t.Root == nil {
		t.Root = l
		return nil
	}

	node := t.Root
	var previous *Leave

	for node != nil {
		previous = node
		if l.Data <= node.Data {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	if l.Data <= previous.Data {
		previous.Left = l
	} else {
		previous.Right = l
	}

	return nil
}

func (l *Leave) InOrder() {
	if l == nil {
		return
	}

	l.Left.InOrder()
	fmt.Print(l.Data)
	l.Right.InOrder()
}

func (l *Leave) PreOrder() {
	if l == nil {
		return
	}

	fmt.Print(l.Data)
	l.Left.PreOrder()
	l.Right.PreOrder()
}

func (l *Leave) PostOrder() {
	if l == nil {
		return
	}

	l.Left.PostOrder()
	l.Right.PostOrder()
	fmt.Print(l.Data)
}
