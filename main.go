package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"ohitsjoe.co/gogorithms/algorithms"
	"ohitsjoe.co/gogorithms/models"
	"os"
)

func main() {
	file, err := os.Open("./mazes/4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mazeBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	maze := string(mazeBytes)

	fmt.Printf("\n\nmaze \n%v\n\n", maze)
	graph := models.NewGraph(maze)

	algorithms.DFSIterative(&graph)
}


