package main

func main() {
	f := PathFinder{
		gridCol: 10,
		gridRow: 10,
	}
	f.Init()
	f.SetStartNode(0, 0)
	f.SetEndNode(8, 6)
	f.Search()
}
