package Node

import "fmt"

type Node struct {
	Data string
}

type Edge struct {
	To   *Node
	Cost float64
}

func (n *Node) ToString() string {
	return fmt.Sprintf("%v", n.Data)
}
