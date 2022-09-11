package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	// empty
	if len(q.items) == 0 {
		return -1
	}
	item, items := q.items[0], q.items[1:]
	q.items = items
	return item
}

func main() {
	q := Queue{}
	fmt.Println("=======Enqueue=======")
	q.Enqueue(1)
	fmt.Println(q)
	q.Enqueue(2)
	fmt.Println(q)
	q.Enqueue(3)
	fmt.Println(q)
	fmt.Println("=======Dequeue=======")
	println(q.Dequeue())
	fmt.Println(q)
	println(q.Dequeue())
	fmt.Println(q)
	println(q.Dequeue())
	fmt.Println(q)
}
