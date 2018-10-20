package main

import (
	"fmt"
	. "graphs/Graph"
	. "graphs/Node"
)

func main() {
	println("App running.")

	var graph = new(Graph)

	var nodeA = Node{Data: "A"}
	var nodeB = Node{Data: "B"}
	var nodeC = Node{Data: "C"}
	var nodeD = Node{Data: "D"}
	var nodeE = Node{Data: "E"}
	var nodeF = Node{Data: "F"}

	graph.AddVertex(&nodeA)
	graph.AddVertex(&nodeB)
	graph.AddVertex(&nodeC)
	graph.AddVertex(&nodeD)
	graph.AddVertex(&nodeE)
	graph.AddVertex(&nodeF)

	graph.AddEdge(&nodeA, &nodeB)
	graph.AddEdge(&nodeB, &nodeC)
	graph.AddEdge(&nodeC, &nodeF)
	graph.AddEdge(&nodeA, &nodeE)
	graph.AddEdge(&nodeF, &nodeD)

	var nodes, _ = graph.DFS(&nodeA, nil)
	for _, node := range nodes {
		fmt.Printf("%v ", node.ToString())
	}

	var _, isPath = graph.DFS(&nodeB, &nodeE)
	fmt.Println("\nPath from B to E: %v", isPath)
}
