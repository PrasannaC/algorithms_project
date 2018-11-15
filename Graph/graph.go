package Graph

import (
	"fmt"
	. "graphs/Node"
)

type EdgeType int32

const (
	UNIDIRECTIONAL = iota
	BIDIRECTIONAL  = iota
)

type WeightedGraph struct {
	Vertices  Nodes
	Edges     map[Node]Edges
	TotalCost float64
}

func (g *WeightedGraph) AddVertex(node *Node) {
	g.Vertices = append(g.Vertices, node)
}

func (g *WeightedGraph) AddEdge(u, v *Node, cost float64, edgeType EdgeType) {
	if g.Edges == nil {
		g.Edges = make(map[Node]Edges)
	}
	g.Edges[*u] = append(g.Edges[*u], &Edge{From: u, To: v, Cost: cost})
	if edgeType == BIDIRECTIONAL {
		g.Edges[*v] = append(g.Edges[*v], &Edge{From: v, To: u, Cost: cost})
	}
	g.TotalCost += cost
}

func (g *WeightedGraph) FindCost(u, v Node) float64 {
	edges := g.Edges[u]
	for _, val := range edges {
		if val.To.Data == v.Data {
			return val.Cost
		}
	}
	return 0.0
}

func (g *WeightedGraph) ToString() string {
	/*
		Logic:
			for every key in map
				get array of neighbours for every node (key)
				for every node in neighbours
		 			print the data
				print newline
	*/
	s := ""
	for i := 0; i < len(g.Edges); i++ {
		s += fmt.Sprintf("%v ->", g.Vertices[i].ToString())
		var adjacentEdges = g.Edges[*g.Vertices[i]]
		for j := 0; j < len(adjacentEdges); j++ {
			s += fmt.Sprintf(" %v", adjacentEdges[j].To.ToString())
		}
		s += fmt.Sprintln("")
	}
	return s
}

/*
	Public method for DFS on graph.
	g -> Pointer to the instance of graph.
	srcNode -> Pointer to the source node of graph.
	nodeToSearch -> Pointer to the node to which the path must be searched from srcNode.
	Returns: Array of pointers to a node, boolean if nodeToSearch is visited.
*/
func (g *WeightedGraph) DFS(srcNode *Node, nodeToSearch *Node) (Nodes, bool) {
	/*
		if sufficient data is available, go ahead, else return nil
		visited <- Create an empty slice(list) to be filled with nodes in DFS order
		returnList <- Create a map to maintain which nodes have been visited
		call g.recursiveDFS(srcNode,visited,returnList)
		return returnList
	*/

	if g.Vertices == nil || g.Edges == nil {
		return nil, false
	}

	var returnList = make(Nodes, 0)
	//
	var visitedMap = make(map[Node]bool)
	//
	isPresent := g.recursiveDFS(srcNode, &visitedMap, &returnList, nodeToSearch)
	return returnList, isPresent
}

/*
	Recursive DFS method.
	g -> Pointer to the instance of graph.
	srcNode -> Pointer to the source node of graph.
	visited -> Pointer to the map containing boolean values to maintain if node is visited.
	returnList -> Pointer to the list containing nodes in DFS order.
	nodeToSearch -> Pointer to the node to which the path must be searched from srcNode.
	Returns: Boolean if nodeToSearch is visited.
*/
func (g *WeightedGraph) recursiveDFS(srcNode *Node, visited *map[Node]bool, returnList *Nodes, nodeToSearch *Node) bool {
	/*
		mark srcNode as visited
		for every edge (srcNode <-> adjacentNode) do:
			if visited[adjacentNode] == false
			{
				call g.recursiveDFS(srcNode,visited,returnList)
			}
	*/
	_, isKeyPresent := (*visited)[*srcNode]
	if !isKeyPresent {
		(*visited)[*srcNode] = true
		*returnList = append(*returnList, srcNode)
	}
	var adjacentEdges = g.Edges[*srcNode]
	if len(adjacentEdges) > 0 {
		for _, adjacentNode := range adjacentEdges {
			if !(*visited)[*adjacentNode.To] {
				g.recursiveDFS(adjacentNode.To, visited, returnList, nodeToSearch)
			}
		}
	}
	if nodeToSearch == nil {
		return false
	}
	return (*visited)[*nodeToSearch]
}
