package algorithms

import (
	"fmt"
	"ohitsjoe.co/gogorithms/models"
	"time"
)

func DFSIterative(graph *models.Graph){

	queue := []*models.Node{
		graph.Root,
	}

	var foundTheGoal *models.Node

	for len(queue) != 0 {
		lastNode := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if lastNode.Checked == true {
			continue
		}

		if lastNode.Content == "G" {
			foundTheGoal = lastNode
			break
		}

		connectedNodes := make([]*models.Node, len(lastNode.Edges))
		for i, edge := range lastNode.Edges {
			connectedNodes[i] = edge.Destination
		}

		queue = append(queue, connectedNodes...)
		lastNode.Checked = true

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

