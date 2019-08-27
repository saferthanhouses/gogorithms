package models

type Edge struct {
	Weight int
	Origin *Node
	Destination *Node
}