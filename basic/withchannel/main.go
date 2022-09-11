package main

import "fmt"

type Queue struct {
	items chan int
}

func (q *Queue) Enqueue(i int) {
	q.items <- i
}

func (q *Queue) Dequeue() int {
	return <-q.items
}

func main() {
	q := Queue{
		items: make(chan int, 16),
	}
	fmt.Println("=======Enqueue=======")
	q.Enqueue(1)
	// fmt.Println(q)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println("=======Dequeue=======")
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
}
