package main

// imagine that linked list is train
// buat node. in the train node is gerbong
type Node struct {
	prev *Node
	next *Node
	data interface{}
}

// buat list. in the train list is the real train ()
