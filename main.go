package main

import (
	"fmt"
	"graphs/DisjointSets"
	"graphs/FileUtils"
	"graphs/Graph"
	"graphs/MST"
	"graphs/Node"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	graph := CreateGraphFromArguments(arguments)
	fmt.Println("MST using DFS is: ")
	start := time.Now()
	_, cost := MST.KruskalDfs(graph)
	elapsed := time.Since(start)
	fmt.Printf("Total Cost: %v\nTime taken: %v\n", cost, elapsed)
	//fmt.Println(mst.ToString())
	fmt.Println("MST using Union Find is: ")
	start = time.Now()
	_, cost = MST.KruskalUnionFind(graph, false, DisjointSets.BY_RANK)
	elapsed = time.Since(start)
	fmt.Printf("Total Cost: %v\nTime taken: %v\n", cost, elapsed)
	//fmt.Println(mst.ToString())

	fmt.Println("MST using Prim is: ")
	start = time.Now()
	_, cost = MST.Prim(graph)
	elapsed = time.Since(start)
	fmt.Printf("Total Cost: %v\nTime taken: %v\n", cost, elapsed)
	//fmt.Println(mst.ToString())
}

func ParseArguments(args []string) string {
	inputFilePath := ""

	for i, element := range args {
		if element == "-i" || element == "-I" {
			if i == len(args)-1 {
				fmt.Println("file path is not given as next argument")
			} else {
				return args[i+1]
			}
		}
	}

	return inputFilePath
}

func CreateGraphFromArguments(arguments []string) *Graph.WeightedGraph {
	inputFile := ParseArguments(arguments)
	data := FileUtils.ReadFile(inputFile)
	nodeMap := make(map[interface{}]Node.Node, 0)
	var graph = new(Graph.WeightedGraph)

	for _, line := range data {
		lineArr := strings.Split(line, " ")
		if len(lineArr) != 3 {
			continue
		}
		fromData := lineArr[0]
		toData := lineArr[1]
		cost := lineArr[2]
		var node1 Node.Node
		var node2 Node.Node
		node, ok := nodeMap[fromData]
		if !ok {
			node1 = Node.Node{Data: fromData}
			nodeMap[fromData] = node1
			graph.AddVertex(&node1)
		} else {
			node1 = node
		}

		node, ok = nodeMap[toData]
		if !ok {
			node2 = Node.Node{Data: toData}
			nodeMap[toData] = node2
			graph.AddVertex(&node2)
		} else {
			node2 = node
		}

		costVal, err := strconv.ParseFloat(cost, 64)
		if err != nil {
			panic(err)
		}
		graph.AddEdge(&node1, &node2, costVal, Graph.BIDIRECTIONAL)
	}

	return graph
}
