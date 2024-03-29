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
	Run(arguments)
}

func ParseArguments(args []string) (string, string, bool, bool) {
	inputFilePath := ""
	mstType := ""
	enablePC := false
	sortedPrim := false

	for i, element := range args {
		if element == "-i" || element == "-I" {
			if i == len(args)-1 {
				fmt.Println("file path is not given as next argument")
			} else {
				inputFilePath = args[i+1]
			}
		}
		if element == "--kruskal-dfs" || element == "-KD" {
			mstType = "KRUSKAL_DFS"
		}
		if element == "--kruskal-union-find" || element == "-KUF" {
			mstType = "KRUSKAL_UF"
		}
		if element == "--prim" || element == "-P" {
			mstType = "PRIM"
		}
		if element == "--path-compression" || element == "-PC" {
			enablePC = true
		}
		if element == "--sorted" || element == "-S" {
			sortedPrim = true
		}

	}

	return inputFilePath, mstType, enablePC, sortedPrim
}

func Run(arguments []string) {
	inputFile, mstType, enablePC, sortedPrim := ParseArguments(arguments)
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

	if sortedPrim {
		graph.SortEdges()
	}

	var output string
	var outputFileName string
	switch mstType {
	case "KRUSKAL_DFS":
		outputFileName = mstType + "_"
		output = fmt.Sprintf("MST using DFS is: \n")
		start := time.Now()
		_, cost := MST.KruskalDfs(graph)
		elapsed := time.Since(start)
		output = fmt.Sprintf("Total Cost: %v\nTime taken: %v\n", cost, elapsed)
		break
	case "KRUSKAL_UF":
		outputFileName = mstType + "_"
		output = fmt.Sprintf("MST using Union Find is: \n")
		start := time.Now()
		_, cost := MST.KruskalUnionFind(graph, enablePC, DisjointSets.BY_RANK)
		elapsed := time.Since(start)
		output = fmt.Sprintf("Total Cost: %v\nTime taken: %v\n", cost, elapsed)

		break
	case "PRIM":
		var sorted_name = "UNSORTED"
		if sortedPrim {
			sorted_name = "SORTED"
		}
		outputFileName = mstType + "_" + sorted_name + "_"
		output = fmt.Sprintf("MST using Prim is: \n")
		start := time.Now()
		_, cost := MST.Prim(graph)
		elapsed := time.Since(start)
		output = fmt.Sprintf("Total Cost: %v\nTime taken: %v\n", cost, elapsed)
		break
	default:
		outputFileName = "TSP_PRIM_"
		start := time.Now()
		mst, _ := MST.Prim(graph)
		_, cost := Graph.TSP(graph, mst)
		elapsed := time.Since(start)
		output = fmt.Sprintf("Total Cost: %v\nTime taken: %v\n", cost, elapsed)
	}

	arr := strings.Split(inputFile, "/")
	t := arr[len(arr)-1]
	file, err := os.Create(outputFileName + t)
	if err != nil {
		panic(err)
	}
	file.WriteString(output)
	file.Close()
}
