package main

import (
	"fmt"
	"graphs/DisjointSets"
	. "graphs/Graph"
	"graphs/MST"
	. "graphs/Node"
)

func main() {
	println("App running.")

	var graph = new(WeightedGraph)

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

	graph.AddEdge(&nodeA, &nodeB, 10, BIDIRECTIONAL)
	graph.AddEdge(&nodeE, &nodeC, 20, BIDIRECTIONAL)
	graph.AddEdge(&nodeA, &nodeE, 30, BIDIRECTIONAL)
	graph.AddEdge(&nodeC, &nodeB, 40, BIDIRECTIONAL)
	graph.AddEdge(&nodeA, &nodeF, 50, BIDIRECTIONAL)
	graph.AddEdge(&nodeF, &nodeE, 60, BIDIRECTIONAL)
	graph.AddEdge(&nodeE, &nodeD, 70, BIDIRECTIONAL)
	graph.AddEdge(&nodeD, &nodeC, 80, BIDIRECTIONAL)
	graph.AddEdge(&nodeB, &nodeE, 90, BIDIRECTIONAL)

	fmt.Println(graph.ToString())

	fmt.Println("DFS Sequence: ")
	var nodes, _ = graph.DFS(&nodeA, nil)
	for _, node := range nodes {
		fmt.Printf("%v ", node.ToString())
	}

	var _, isPath = graph.DFS(&nodeB, &nodeE)
	fmt.Printf("\n\nPath from B to E: %v\n\n", isPath)

	fmt.Println("MST using DFS is: ")
	mst := MST.KruskalDfs(graph)
	fmt.Println(mst.ToString())
	fmt.Println("MST using Union Find is: ")
	mst = MST.KruskalUnionFind(graph, false, DisjointSets.BY_RANK)
	fmt.Println(mst.ToString())

	d := DisjointSets.CreateDisjointSet(true, DisjointSets.BY_SIZE)
	d.MakeSet(10)
	d.MakeSet(20)
	d.Union(10, 20)

	fmt.Println("MST using Prims is: ")
	mst = MST.Prims(graph)
	fmt.Println(mst.ToString())
}
