package queue

import "fmt"

type Queue struct {
	queue []int
	tail  int
}

func (q *Queue) Enqueue(value int) {
	q.queue = append([]int{value}, q.queue...)
	q.tail = q.tail + 1
}

func (q *Queue) Dequeue() {
	if len(q.queue) <= 0 {
		fmt.Println("Queue is empty")
	} else {

		q.queue = q.queue[1:]
		q.tail = q.tail - 1
	}
}

func QueueDS() {

	var queue Queue = Queue{}

	queue.Enqueue(3)
	queue.Enqueue(6)
	queue.Enqueue(10)
	queue.Enqueue(8)

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	fmt.Println(queue)

}
