package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

func (n *Node) GetNode() *Node {
	return n
}

func (n *Node) GetData() int {
	return n.data
}

func (n *Node) ReturnNext() *Node {
	return n.next
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) AddNode(value int) {
	var newNode = &Node{
		data: value,
		next: nil,
	}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode

}

func (l *LinkedList) Traverse() {
	var curr = l.head

	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}

}
func traverseFromHead(head *Node) {
	var curr = head

	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}

}

func (l *LinkedList) DeleteNode(value int) {
	curr := l.head

	if curr == nil {
		return
	}

	if curr.data == value {
		curr = nil
		return
	}

	prev := l.head

	for curr != nil {

		if curr.data == value {
			prev.next = curr.next
			break
		}

		prev = curr
		curr = curr.next
	}

}

func (l *LinkedList) UpdateNode(value int, newValue int) {

	curr := l.head

	if curr == nil {
		return
	}

	if curr.data == value {
		curr.data = newValue
		return
	}

	for curr != nil {

		if curr.data == value {
			curr.data = newValue
			break
		}

		curr = curr.next

	}

}

func (l *LinkedList) Reverse() *Node {

	curr := l.head

	var prev *Node

	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}

	return prev
}

func LinkedListDSA() {

	var linkedList LinkedList = LinkedList{}

	linkedList.AddNode(2)
	linkedList.AddNode(3)
	linkedList.AddNode(5)
	linkedList.AddNode(6)

	fmt.Println("End")

	// linkedList.deleteNode(5)

	// linkedList.updateNode(5, 4)

	linkedList.Traverse()
	head := linkedList.Reverse()
	fmt.Println("head", head)
	traverseFromHead(head)

}
