package models

type Node struct {
	Edges   []*Edge
	Content string
	Checked bool
}
