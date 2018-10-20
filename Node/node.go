package Node

import "fmt"

type Node struct {
	Data string
}

func (n *Node) ToString() string {
	return fmt.Sprintf("%v", n.Data)
}
