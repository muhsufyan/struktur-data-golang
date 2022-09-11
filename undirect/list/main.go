package main

import "fmt"

/*
   Go Program for
   Undirected graph representation by using adjacency list
*/
type AjlistNode struct {
	// Vertices node key
	id   int
	next *AjlistNode
}

func getAjlistNode(id int) *AjlistNode {
	var me *AjlistNode = &AjlistNode{}
	// Set value of node key
	me.id = id
	me.next = nil
	return me
}

type Vertices struct {
	data int
	next *AjlistNode
	last *AjlistNode
}

func getVertices(data int) *Vertices {
	var me *Vertices = &Vertices{}
	me.data = data
	me.next = nil
	me.last = nil
	return me
}

type Graph struct {
	// Number of Vertices
	size int
	node []*Vertices
}

func getGraph(size int) *Graph {
	var me *Graph = &Graph{size, make([]*Vertices, size)}
	me.setData()
	return me
}

// Set initial node value
func (this *Graph) setData() {
	if this.size <= 0 {
		fmt.Println("\nEmpty Graph")
	} else {
		for index := 0; index < this.size; index++ {
			this.node[index] = getVertices(index)
		}
	}
}

// Connect two nodes
func (this *Graph) connect(start, last int) {
	var edge *AjlistNode = getAjlistNode(last)
	if this.node[start].next == nil {
		this.node[start].next = edge
	} else {
		// Add edge at the end
		this.node[start].last.next = edge
	}
	// Get last edge
	this.node[start].last = edge
}

//  Handling the request of adding new edge
func (this *Graph) addEdge(start, last int) {
	if start >= 0 && start < this.size && last >= 0 && last < this.size {
		this.connect(start, last)
		if start == last {
			// When self loop
			return
		}
		this.connect(last, start)
	} else {
		// When invalid nodes
		fmt.Println("\nHere Something Wrong")
	}
}
func (this Graph) printGraph() {
	if this.size > 0 {
		// Print graph ajlist Node value
		for index := 0; index < this.size; index++ {
			fmt.Print("\nAdjacency list of vertex ", index, " :")
			var temp *AjlistNode = this.node[index].next
			for temp != nil {
				// Display graph node
				fmt.Print("  ", this.node[temp.id].data)
				// Visit to next edge
				temp = temp.next
			}
		}
	}
}
func main() {
	// 5 implies the number of nodes in graph
	var g *Graph = getGraph(5)
	// Connect node with an edge
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(0, 4)
	g.addEdge(1, 2)
	g.addEdge(1, 4)
	g.addEdge(2, 3)
	g.addEdge(3, 4)
	// print graph element
	g.printGraph()
}
