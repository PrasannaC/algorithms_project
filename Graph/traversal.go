package Graph

import (
	"graphs/Node"
)

func TSP(graph, mst *WeightedGraph) (*WeightedGraph, float64) {
	sourceNode := mst.Vertices[0]
	sourceNodeEdges := mst.Edges[*sourceNode]
	var visitedMap = make(map[Node.Node]bool)
	var tspCycle = new(WeightedGraph)
	var list = make([]Node.Node, 0)
	mst.SortEdges()
	recursivePreorder(mst, &list, sourceNode, nil, sourceNodeEdges, &visitedMap)
	//for _, v := range list {
	//	fmt.Println(v.Data)
	//}

	for i, j := 0, 1; j < len(list); {
		u, v := list[i], list[j]
		if i == 0 {
			tspCycle.AddVertex(&u)
		}
		tspCycle.AddVertex(&v)
		cost := graph.FindCost(u, v)
		tspCycle.AddEdge(&u, &v, cost, BIDIRECTIONAL)
		i++
		j++
	}
	u, v := list[0], list[len(list)-1]
	cost := graph.FindCost(u, v)
	tspCycle.AddEdge(&u, &v, cost, BIDIRECTIONAL)
	//fmt.Println(tspCycle.ToString())
	//fmt.Printf("TSP cost: %v", tspCycle.TotalCost)
	return tspCycle, tspCycle.TotalCost
}

func recursivePreorder(graph *WeightedGraph, tsp *[]Node.Node, source *Node.Node, prev *Node.Node, edges Node.Edges, visited *map[Node.Node]bool) float64 {
	cost := float64(0)
	_, isKeyPresent := (*visited)[*source]
	if !isKeyPresent {
		(*visited)[*source] = true
		*tsp = append(*tsp, *source)
	}

	for _, v := range edges {
		if !(*visited)[*v.To] {
			recursivePreorder(graph, tsp, v.To, source, graph.Edges[*(*v).To], visited)
		}
	}
	return cost
}
