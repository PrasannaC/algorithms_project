package MST

import (
	"container/heap"
	"graphs/Graph"
	"graphs/Node"
	"math"
)

func Prims(graph *Graph.WeightedGraph) *Graph.WeightedGraph {

	queue := make(PriorityQueue, 0)
	count := 0
	MST := Graph.WeightedGraph{}
	//assign priorities to every vertex.
	for _, vertex := range graph.Vertices {
		MST.AddVertex(vertex)
		if count == 0 {
			item := PQItem{value: vertex, priority: 0.0}
			heap.Push(&queue, &item)
			count++
		} else {
			item := PQItem{value: vertex, priority: math.Inf(1)}
			heap.Push(&queue, &item)
		}
	}
	visited := make(map[*Node.Node]bool, 0)
	weights := make(map[*Node.Node]float64, 0)
	parent := make(map[*Node.Node]*Node.Node, 0)
	heap.Init(&queue)
	for queue.Len() > 0 {
		minVertexItem := queue.Pop().(*PQItem).value
		for _, edge := range graph.Edges[*minVertexItem] {
			u := edge.From
			v := edge.To
			if _, ok := visited[v]; !ok {
				visited[v] = true
				weights[v] = edge.Cost
				parent[v] = u
			} else {
				//node is visited.
				if weights[v] > edge.Cost {
					weights[v] = edge.Cost
					parent[v] = u
				}
			}
		}
	}

	//make mst edges
	for key, value := range parent {
		MST.AddEdge(value, key, weights[value], Graph.BIDIRECTIONAL)
	}
	return &MST
}
