package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"ohitsjoe.co/gogorithms/models"
	"os"
)

func main() {
	file, err := os.Open("./mazes/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mazeBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	maze := string(mazeBytes)

	graph := models.NewGraph(maze)

	fmt.Println(graph.Root.Content)
}


