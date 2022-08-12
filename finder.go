package main

import (
	"fmt"
	"math"
	"sort"
)

type PathFinder struct {
	gridRow    int
	gridCol    int
	startNode  *Node
	endNode    *Node
	openNodes  []*Node
	closeNodes []*Node
	grid       *Grid
}

func (f *PathFinder) Init() {
	f.grid = &Grid{
		rows: f.gridRow,
		cols: f.gridCol,
	}
	f.grid.GenGrid()
}

func (f *PathFinder) SetStartNode(x, y int) {
	f.startNode = f.grid.GetNode(x, y)
}

func (f *PathFinder) SetEndNode(x, y int) {
	f.endNode = f.grid.GetNode(x, y)
}

func (f *PathFinder) Search() {
	currentNode := f.startNode

	for !f.endNode.Eq(currentNode) {
		startX := int(math.Max(float64(currentNode.x-1), 0))
		startY := int(math.Max(float64(currentNode.y-1), 0))
		endX := int(math.Min(float64(currentNode.x+1), float64(f.grid.cols-1)))
		endY := int(math.Min(float64(currentNode.y+1), float64(f.grid.rows-1)))

		for i := startX; i <= endX; i++ {
			for j := startY; j <= endY; j++ {
				test := f.grid.GetNode(i, j)
				if currentNode.Eq(test) {
					continue
				}

				cost := 2
				if currentNode.x == test.x || currentNode.y == test.y {
					cost = 1
				}

				costG := currentNode.costG + cost
				costH := f.diagonal(test)
				costF := costG + costH

				if f.isOpen(test) || f.isClose(test) {
					if test.costF > costF {
						test.costF = costF
						test.costG = costG
						test.costH = costH
						test.parentNode = currentNode
					}
				} else {
					test.costF = costF
					test.costG = costG
					test.costH = costH
					test.parentNode = currentNode
					f.openNodes = append(f.openNodes, test)
				}
			}
		}
		f.closeNodes = append(f.closeNodes, currentNode)

		if len(f.openNodes) == 0 {
			fmt.Println("見つからない！")
			return
		}

		sort.Slice(f.openNodes, func(i, j int) bool {
			return f.openNodes[i].costF < f.openNodes[j].costF
		})

		currentNode = f.openNodes[0]
		f.openNodes = f.openNodes[:0+copy(f.openNodes[0:], f.openNodes[1:])]
	}

	f.buildPath()
}

func (f *PathFinder) buildPath() {
	node := f.endNode
	var nodes []*Node
	nodes = append(nodes, node)
	for !node.Eq(f.startNode) {
		node = node.parentNode
		nodes = append(nodes, node)
	}

	fmt.Println("◇ がスタート")
	fmt.Println("◆ がゴール")
	fmt.Println("■ が経路")
	for i := 0; i < f.gridRow; i++ {
	a:
		for j := 0; j < f.gridCol; j++ {
			if f.startNode.x == j && f.startNode.y == i {
				fmt.Print("◇  ")
				continue
			}
			if f.endNode.x == j && f.endNode.y == i {
				fmt.Print("◆  ")
				continue
			}
			for _, n := range nodes {
				if n.y == i && n.x == j {
					fmt.Print("■  ")
					continue a
				}
			}
			fmt.Print("□  ")
		}
		fmt.Println()
	}
}

func (f *PathFinder) diagonal(node *Node) int {
	dx := math.Abs(float64(node.x - f.endNode.x))
	dy := math.Abs(float64(node.y - f.endNode.y))
	diag := math.Min(dx, dy)
	straight := dx + dy

	return int(1*diag + 1*(straight-2*diag))
}

func (f *PathFinder) isOpen(node *Node) bool {
	for _, oNode := range f.openNodes {
		if oNode.Eq(node) {
			return true
		}
	}

	return false
}

func (f *PathFinder) isClose(node *Node) bool {
	for _, cNode := range f.closeNodes {
		if cNode.Eq(node) {
			return true
		}
	}

	return false
}
