package algorithms

import (
	"fmt"
	"ohitsjoe.co/gogorithms/models"
	"time"
)

func BFSIterative(graph *models.Graph){

	queue := []*models.Node{
		graph.Root,
	}

	var foundTheGoal *models.Node

	for len(queue) != 0 {
		firstNode := queue[0]
		queue = queue[1:]

		if firstNode.Checked == true {
			continue
		}

		if firstNode.Content == "G" {
			foundTheGoal = firstNode
			break
		}

		connectedNodes := make([]*models.Node, len(firstNode.Edges))
		for i, edge := range firstNode.Edges {
			connectedNodes[i] = edge.Destination
		}

		queue = append(queue, connectedNodes...)
		firstNode.Checked = true

		graph.PrintGraph()
		time.Sleep(4 * time.Millisecond)
	}

	if foundTheGoal != nil {
		fmt.Println("Found the goal :D")
	} else {
		fmt.Println("No route to the goal ... :(")
	}

	graph.PrintGraph()
}

