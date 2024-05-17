package main

import (
	"fmt"
)

func main() {
	graph := graphFromFile("adjacency_matrix")
	fmt.Println("Graph: ", graph)
	fmt.Println("BFS: ", bfs(graph, 0))
	fmt.Println("Find components via BFS: ", bfsGetComps(graph))
	fmt.Println("DFS: ", dfs(graph, 0))
	fmt.Println("Find strong components via DFS: ", dfsGetComps(graph))
	fmt.Scanf("h")
}
