package main

type Grid struct {
	rows     int
	cols     int
	nodeGrid [][]*Node
}

func (g *Grid) GenGrid() {
	g.nodeGrid = make([][]*Node, g.rows, g.rows)
	for i := 0; i < g.rows; i++ {
		g.nodeGrid[i] = make([]*Node, g.cols, g.cols)
		for j := 0; j < g.cols; j++ {
			g.nodeGrid[i][j] = &Node{
				x: j,
				y: i,
			}
		}
	}
}

func (g *Grid) GetNode(x, y int) *Node {
	return g.nodeGrid[y][x]
}
