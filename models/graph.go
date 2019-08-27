package models

import (
	"strconv"
	"strings"
)

type Graph struct {
	Root *Node
}

func getNodeKey(a int, b int) string {
	name := strings.Builder{}
	name.WriteString(strconv.Itoa(a))
	name.WriteString("-")
	name.WriteString(strconv.Itoa(b))
	return name.String()
}

func getOrMakeNode(a int, b int, mazeRows []string, nodeMap map[string]*Node) *Node {
	key := getNodeKey(a, b)
	_, ok := nodeMap[key]
	if !ok {
		content := string(mazeRows[a][b])
		node :=  Node{
			[]*Edge{},
			content,
			false,
		}
		nodeMap[key] = &node
	}

	return nodeMap[key]
}

func createEdgeAndAppend(sourceNode *Node, targetNode *Node){
	edge := Edge{
		1,
		sourceNode,
		targetNode,
	}

	sourceNode.Edges = append(sourceNode.Edges, &edge)
}


func NewGraph(mazeTextInput string) Graph {
	mazeRows := strings.Split(mazeTextInput, "\n")

	nodeMap := make(map[string]*Node)
	graph := Graph{}

	// Process the text maze to produce a series of Nodes & Edges
	for i := 0; i < len(mazeRows); i++ {
		row := mazeRows[i]
		for j := 0; j < len(row); j ++ {
			node := getOrMakeNode(i, j, mazeRows, nodeMap)
			if node.Content == "@" {
				graph.Root = node
			}

			if i-1 >= 0 && contentIsNode(mazeRows[i-1][j]) {
				above := getOrMakeNode(i-1, j, mazeRows, nodeMap)
				createEdgeAndAppend(node, above)
			}

			if i+1 < len(mazeRows) && contentIsNode(mazeRows[i+1][j]) {
				below := getOrMakeNode(i+1, j, mazeRows, nodeMap)
				createEdgeAndAppend(node, below)
			}

			if j-1 >= 0 && contentIsNode(mazeRows[i][j-1]) {
				left := getOrMakeNode(i, j-1, mazeRows, nodeMap)
				createEdgeAndAppend(node, left)
			}

			if j+1 < len(row) && contentIsNode(mazeRows[i][j+1]) {
				right := getOrMakeNode(i, j+1, mazeRows, nodeMap)
				createEdgeAndAppend(node, right)
			}

		}
	}

	return graph
}

func contentIsNode(u uint8) bool {
	return u == 120 || u == 71 || u == 64
}


