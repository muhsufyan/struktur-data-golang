package main

// https://www.youtube.com/watch?v=TyA1M0L0EM4
import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

// Binary Tree Traversal
// Preorder Value - Left - Right
// Inorder Left - Value - Right
// Postorder Left - Right - Value
// pre, post, and in are references to the print

func main() {
	base := initBTree()
	fmt.Println("pre order")
	base.preOrder()
	fmt.Println("in order")
	base.inOrder()
	fmt.Println("post order")
	base.postOrder()
}

func initBTree() *Node {
	five := Node{val: 5}
	three := Node{val: 3, left: &five}
	six := Node{val: 6}
	fifteen := Node{val: 15, left: &three, right: &six}
	nine := Node{val: 9}
	eight := Node{val: 8}
	two := Node{val: 2, left: &nine, right: &eight}
	thirty := Node{val: 30, right: &two}
	return &Node{val: 10, right: &thirty, left: &fifteen}
}

func (n *Node) preOrder() {
	if n != nil {
		fmt.Println(n.val)
		n.left.preOrder()
		n.right.preOrder()
	}
}

func (n *Node) inOrder() {
	if n != nil {
		n.left.inOrder()
		fmt.Println(n.val)
		n.right.inOrder()
	}
}

func (n *Node) postOrder() {
	if n != nil {
		n.left.postOrder()
		n.right.postOrder()
		fmt.Println(n.val)
	}
}
