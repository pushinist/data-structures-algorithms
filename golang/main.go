package main

import (
	"fmt"
)

func main() {
	graph := graphFromFile("adjacency_matrix")
	fmt.Println(graph)
	fmt.Println(bfs(graph, 0))
	fmt.Println(bfsGetComps(graph))
	fmt.Println(dfs(graph, 0))
	fmt.Println(dfsGetComps(graph))
	fmt.Scanf("h")
}
