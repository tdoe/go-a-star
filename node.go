package main

type Node struct {
	x          int
	y          int
	costF      int
	costG      int
	costH      int
	parentNode *Node
}

func (n *Node) Eq(a *Node) bool {
	return a.y == n.y && a.x == n.x
}
