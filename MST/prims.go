package MST

import (
	"container/heap"
	"graphs/Graph"
	"graphs/Node"
)

/*
	Returns the minimum spanning tree for a graph
	using Prim's algorithm.
	graph -> Pointer to the instance of the input graph.
*/
func Prim(graph *Graph.WeightedGraph) *Graph.WeightedGraph {
	queue := make(PriorityQueue, len(graph.Vertices))
	vertexPQMap := make(map[Node.Node]*PQItem, 0)
	count := 0
	MST := Graph.WeightedGraph{}
	//assign priorities to every vertex.
	for i, vertex := range graph.Vertices {
		MST.AddVertex(vertex)
		var item PQItem
		if count == 0 {
			item = PQItem{value: *vertex, priority: float64(0)}
			queue[i] = &item
			count++
		} else {
			item = PQItem{value: *vertex, priority: float64(999999999)}
			queue[i] = &item
		}
		queue[i].Index = i
		vertexPQMap[*vertex] = &item
	}
	heap.Init(&queue)

	visited := make(map[Node.Node]bool, 0)
	parent := make(map[*Node.Node]*Node.Node, 0)

	for queue.Len() > 0 {
		minVertexItem := queue.Pop().(*PQItem)
		visited[minVertexItem.value] = true
		for _, edge := range graph.Edges[minVertexItem.value] {
			u := edge.From
			v := edge.To
			if isVisited, _ := visited[*v]; !isVisited {
				visited[*v] = true
				queue.update(vertexPQMap[*v], vertexPQMap[*v].value, edge.Cost)
				parent[v] = u
			} else {
				//node is visited.
				if parent[u] != v && vertexPQMap[*v].priority > edge.Cost {
					queue.update(vertexPQMap[*v], vertexPQMap[*v].value, edge.Cost)
					parent[v] = u
				}
			}
		}
	}

	//make mst edges
	for key, value := range parent {
		MST.AddEdge(value, key, vertexPQMap[*value].priority, Graph.BIDIRECTIONAL)
	}
	return &MST
}
