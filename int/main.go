package main

import "fmt"

// menyimpan data
type Stack struct {
	// if want accept all data type, the data type is interface (no integer in this case)
	items []int
}

// memasukkan data (append atau add)
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// mengambil data yang mana akan diambil pada data yg paling terakhir dimasukkan (delete)
func (s *Stack) Pop() int {
	left := len(s.items)
	// cek if empty
	if left == 0 {
		return -1
	}

	item, items := s.items[left-1], s.items[0:left-1]
	s.items = items
	return item
}

func main() {
	s := Stack{}
	println("Push (append / add)")
	s.Push(1)
	fmt.Println(s)
	s.Push(2)
	fmt.Println(s)
	s.Push(3)
	fmt.Println(s)
	s.Push(4)
	fmt.Println(s)
	println("Pop (delete)")
	println(s.Pop())
	fmt.Println(s)
	println(s.Pop())
	fmt.Println(s)
	println(s.Pop())
	fmt.Println(s)
	println(s.Pop())
	fmt.Println(s)
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
}
