package algorithms

import (
	"fmt"
	"ohitsjoe.co/gogorithms/models"
)

func DFSIterative(graph *models.Graph){

	queue := []*models.Node{
		graph.Root,
	}

	var foundTheGoal bool = false

	for foundTheGoal == false {
		lastNode := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if lastNode.Checked == true {
			continue
		}

		if lastNode.Content == "G" {
			fmt.Println("Found the goal")
			foundTheGoal = true
			break
		}

		connectedNodes := make([]*models.Node, len(lastNode.Edges))
		for i, edge := range lastNode.Edges {
			connectedNodes[i] = edge.Destination
		}

		//fmt.Printf("node.Edges %v\n", node.Edges)
		queue = append(queue, connectedNodes...)
		//fmt.Printf("queue: %v\n", queue)
		lastNode.Checked = true
	}
}

