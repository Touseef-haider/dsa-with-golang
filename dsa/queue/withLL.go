package queue

import "fmt"

type Node struct {
	data int
	next *Node
}

type QueueWithLinkedlist struct {
	front *Node
	rear  *Node
}

func (q *QueueWithLinkedlist) enqueue(value int) {

	newNode := &Node{
		data: value,
	}

	if q.front == nil {
		q.front = newNode
		q.rear = newNode

		return
	}

	q.rear.next = newNode
	q.rear = newNode

}

func (q *QueueWithLinkedlist) dequeue() *Node {
	temp := q.front

	if q.front == nil {
		return q.front
	}

	q.front = temp.next

	return temp

}

func (q *QueueWithLinkedlist) print() {
	curr := q.front

	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}

}

func QueueLinkedList() {
	var q QueueWithLinkedlist = QueueWithLinkedlist{}

	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)

	q.dequeue()

	q.print()
}
