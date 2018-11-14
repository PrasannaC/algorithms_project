package Node

import "fmt"

type Node struct {
	Data string
}

type Nodes []*Node

func (n *Node) ToString() string {
	return fmt.Sprintf("%v", n.Data)
}

type Edge struct {
	From *Node
	To   *Node
	Cost float64
}

//Creating wrapper type Edges to implement sort interface
type Edges []*Edge

func (e Edges) Len() int {
	return len(e)
}
func (e Edges) Swap(i, j int) {
	(e)[i], (e)[j] = (e)[j], (e)[i]
}
func (e Edges) Less(i, j int) bool {
	return (e)[i].Cost < (e)[j].Cost
}

// An PQItem is something we store in a priority queue.
type PQItem struct {
	Value    Node    // The value of the item; arbitrary.
	Priority float64 // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}
