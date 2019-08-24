package main

import "strings"

type Graph struct {
	rootNode Node
}

type Edge struct {
	weight int
}

func NewGraph(mazeTextInput string) Graph {
	mazeRows := strings.Split(mazeTextInput, "\n")
	for i := 0; i < len(mazeRows); i++ {

	}
}

type Node struct {

}
