package Graph

import (
	"fmt"
	. "graphs/Node"
)

type Graph struct {
	Vertices []*Node
	Edges    map[Node][]*Node
}

func (g *Graph) AddVertex(node *Node) {
	g.Vertices = append(g.Vertices, node)
}

func (g *Graph) AddEdge(u, v *Node) {
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}
	g.Edges[*u] = append(g.Edges[*u], v)
	g.Edges[*v] = append(g.Edges[*v], u)
}

func (g *Graph) ToString() {
	/*
	Logic:
		for every key in map
			get array of neighbours for every node (key)
			for every node in neighbours
	 			print it out in a line
			print newline
	*/

	for i := 0; i <= len(g.Edges); i++ {
		fmt.Printf("%v ->", g.Vertices[i].ToString())
		var neighbours = g.Edges[*g.Vertices[i]]
		for j := 0; j <= len(neighbours); j++ {
			fmt.Printf(" %v", neighbours[j].ToString())
		}
		fmt.Println("")
	}
}

/*
	Public method for DFS on graph.
	g -> Pointer to the instance of graph.
	srcNode -> Pointer to the source node of graph.
	nodeToSearch -> Pointer to the node to which the path must be searched from srcNode.
	Returns: Array of pointers to a node, boolean if nodeToSearch is visited.
 */
func (g *Graph) DFS(srcNode *Node, nodeToSearch *Node) ([]*Node, bool) {
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

	var returnList = make([]*Node, 0)
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
func (g *Graph) recursiveDFS(srcNode *Node, visited *map[Node]bool, returnList *[]*Node, nodeToSearch *Node) bool {
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
			if !(*visited)[*adjacentNode] {
				g.recursiveDFS(adjacentNode, visited, returnList, nodeToSearch)
			}
		}
	}
	if nodeToSearch == nil {
		return false
	}
	return (*visited)[*nodeToSearch]
}
