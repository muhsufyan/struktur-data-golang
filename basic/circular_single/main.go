package main

import (
	"fmt"
	"math/rand"
)

type Nodecircularsingle struct {
	info interface{}
	next *Nodecircularsingle
}

type Listcircularsingle struct {
	head *Nodecircularsingle
}

func (l *Listcircularsingle) Insertcircularsingle(d interface{}) {
	list := &Nodecircularsingle{info: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = list
	}
}
func ShowCircular(l *Listcircularsingle) {
	p := l.head
	for {
		if p.next == l.head {
			fmt.Printf("-> %v ", p.info)
			break
		}
		fmt.Printf("-> %v ", p.info)
		p = p.next
	}
}

func ConvertSinglyToCircular(l *Listcircularsingle) {
	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = l.head
}
func main() {
	sl := Listcircularsingle{}
	for i := 0; i < 5; i++ {
		sl.Insertcircularsingle(rand.Intn(100))
	}
	ConvertSinglyToCircular(&sl)
	fmt.Println()
	ShowCircular(&sl)
}
