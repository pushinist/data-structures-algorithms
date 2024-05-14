package main

import (
	"fmt"
)

func main() {
	graph := [][]uint{
		{1, 2}, // 0
		{0},    // 1
		{0},    // 2
		{},     // 3
	}
	fmt.Println(bfs(graph, 0))
	fmt.Println(getComps(graph))
	fmt.Scanf("h")
}
