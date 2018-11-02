package MST

import (
	"graphs/DisjointSets"
	. "graphs/Graph"
	. "graphs/Node"
	"sort"
)

/*
	Returns the minimum spanning tree for a graph
	using Kruskal's algorithm which is internally using
	DFS to check for cycles.
	g -> Pointer to the instance of the input graph.
*/
func KruskalDfs(g *WeightedGraph) *WeightedGraph {
	/*
		MST <- Initialize empty MST with g.Vertices and {} Edges
		sortedEdges <- Sort Edges in order of increasing Cost
		for every edge (u <-> v) in sortedEdges do:
			if not DFS(MST,u,v)
				MST.Edges = MST.Edges U {edge}
		return MST
	*/
	if g == nil {
		return nil
	}
	MST := WeightedGraph{}
	for _, node := range g.Vertices {
		//not adding by reference. just to keep MST and graph as disjoint
		var copyNode = *node
		MST.AddVertex(&copyNode)
	}

	var edgesList = make(Edges, 0)
	for _, value := range g.Edges {
		for _, edge := range value {
			edgesList = append(edgesList, edge)
		}
	}
	sort.Sort(edgesList)
	for _, edge := range edgesList {
		if _, ok := MST.DFS(edge.From, edge.To); !ok {
			//copying to not add by reference
			var from = *(edge.From)
			var to = *(edge.To)
			MST.AddEdge(&from, &to, edge.Cost, BIDIRECTIONAL)
		}
	}
	return &MST
}

func KruskalUnionFind(g *WeightedGraph, usePathCompression bool, unionBy DisjointSets.UnionByType) *WeightedGraph {

	if g == nil {
		return nil
	}
	MST := WeightedGraph{}
	disjointSet := DisjointSets.CreateDisjointSet(usePathCompression, unionBy)
	for _, node := range g.Vertices {
		//not adding by reference. just to keep MST and graph as disjoint
		var copyNode = *node
		MST.AddVertex(&copyNode)
		disjointSet.MakeSet(copyNode)
	}
	var edgesList = make(Edges, 0)
	for _, value := range g.Edges {
		for _, edge := range value {
			edgesList = append(edgesList, edge)
		}
	}
	sort.Sort(edgesList)
	for _, edge := range edgesList {

		var from = *(edge.From)
		var to = *(edge.To)
		var fromRoot = disjointSet.Find(from)
		var toRoot = disjointSet.Find(to)

		if fromRoot != toRoot {
			MST.AddEdge(&from, &to, edge.Cost, BIDIRECTIONAL)
			disjointSet.Union(from, to)
		}
	}
	return &MST
}
