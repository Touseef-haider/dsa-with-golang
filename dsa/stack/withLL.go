package stack

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type StackWithLL struct {
	head *Node
}

func (l *StackWithLL) push(value int) {

	newNode := &Node{
		data: value,
		next: &Node{},
		prev: &Node{},
	}

	if l.head == nil {
		l.head = newNode
		return
	}

	newNode.prev = l.head

	l.head = newNode

}

func (l *StackWithLL) pop() *Node {

	temp := l.head

	l.head = l.head.prev

	return temp

}

func (l *StackWithLL) print() {
	curr := l.head

	for curr.prev != nil {
		fmt.Println(curr.data)
		curr = curr.prev
	}
}

func StackWithLinkedList() {
	fmt.Println("With linked list")

	s := StackWithLL{}
	s.push(12)
	s.push(13)
	s.push(14)
	s.push(15)

	s.print()

}
