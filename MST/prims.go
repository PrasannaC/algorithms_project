package MST

import (
	pq2 "github.com/jupp0r/go-priority-queue"
	"graphs/Graph"
	"graphs/Node"
)

/*
	Returns the minimum spanning tree for a graph
	using Prim's algorithm.
	graph -> Pointer to the instance of the input graph.
*/
func Prim(graph *Graph.WeightedGraph) (*Graph.WeightedGraph, float64) {
	queue := pq2.New()
	vertexPQMap := make(map[Node.Node]*Node.PQItem, 0)
	count := 0
	MST := Graph.WeightedGraph{}
	//assign priorities to every vertex.
	for _, vertex := range graph.Vertices {
		MST.AddVertex(vertex)
		var item Node.PQItem
		if count == 0 {
			item = Node.PQItem{Value: *vertex, Priority: float64(0)}
			queue.Insert(&item, float64(0))
			count++
		} else {
			item = Node.PQItem{Value: *vertex, Priority: float64(999999999)}
			queue.Insert(&item, float64(999999999))
		}
		vertexPQMap[*vertex] = &item
	}

	//heap.Init(&queue)

	visited := make(map[Node.Node]bool, 0)
	parent := make(map[interface{}]interface{}, 0)

	for queue.Len() > 0 {
		minVertexItem, _ := queue.Pop()
		visited[minVertexItem.(*Node.PQItem).Value] = true
		for _, edge := range graph.Edges[minVertexItem.(*Node.PQItem).Value] {
			u := edge.From
			v := edge.To
			if isVisited, _ := visited[*v]; !isVisited {
				visited[*v] = true
				vertexPQMap[*v].Priority = edge.Cost
				queue.UpdatePriority(vertexPQMap[*v], edge.Cost)
				parent[*v] = *u
			} else {
				//node is visited.
				if parent[*u].(Node.Node).Data != v.Data && vertexPQMap[*v].Priority > edge.Cost {
					vertexPQMap[*v].Priority = edge.Cost
					queue.UpdatePriority(vertexPQMap[*v], edge.Cost)
					parent[*v] = *u
				}
			}
		}
	}

	//make mst edges
	totalCost := float64(0)
	for key, value := range parent {
		var v = key.(Node.Node)
		var u = value.(Node.Node)
		MST.AddEdge(&u, &v, vertexPQMap[v].Priority, Graph.BIDIRECTIONAL)
		totalCost += vertexPQMap[v].Priority
	}
	return &MST, totalCost
}
