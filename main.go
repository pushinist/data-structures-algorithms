package main

import (
	"fmt"
)

func main() {
	graph := graphFromFile("adjacency_matrix")
	fmt.Println(graph)
	fmt.Println(bfs(graph, 0))
	fmt.Println(getComps(graph))
	fmt.Scanf("h")
}
