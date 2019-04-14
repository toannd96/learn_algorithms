package main

import "fmt"

type Queue struct {
	slice []int
}

func (q *Queue) Enqueue(n int) {
	q.slice = append(q.slice, n)
}

func (q *Queue) Dequeue() int {
	removedItem := q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return removedItem
}

func main() {
	newQueue := Queue{}
	newQueue.Enqueue(1)
	newQueue.Enqueue(2)
	newQueue.Enqueue(3)

	fmt.Println(newQueue)

	fmt.Println(newQueue.Dequeue())

	fmt.Println(newQueue)

	fmt.Println(newQueue.Dequeue())

	fmt.Println(newQueue)

	fmt.Println(newQueue.Dequeue())

	fmt.Println(newQueue)
}