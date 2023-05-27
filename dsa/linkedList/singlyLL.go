package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) addNode(value int) {
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

func (l *LinkedList) traverse() {
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

func (l *LinkedList) deleteNode(value int) {
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

func (l *LinkedList) updateNode(value int, newValue int) {

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

func (l *LinkedList) reverse() *Node {

	curr := l.head

	var prev *Node

	fmt.Println(prev)

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

	linkedList.addNode(2)
	linkedList.addNode(3)
	linkedList.addNode(5)
	linkedList.addNode(6)

	fmt.Println("End")

	// linkedList.deleteNode(5)

	// linkedList.updateNode(5, 4)

	linkedList.traverse()
	head := linkedList.reverse()
	fmt.Println("head", head)
	traverseFromHead(head)

}
