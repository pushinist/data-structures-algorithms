package main

func bfs(graph [][]uint, start uint) []uint {
	visited := make([]uint, 0)
	queue := make([]uint, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		visited = append(visited, queue[0])
		for _, v := range graph[visited[len(visited)-1]] {
			if !contains(visited, v) && !contains(queue, v) {
				queue = append(queue, v)
			}
		}
		queue = queue[1:]
	}
	return visited

}

func bfsGetComps(graph [][]uint) [][]uint {
	component := bfs(graph, 0)
	components := make([][]uint, 0)
	visited := make([]uint, 0)
	for _, v := range component {
		if !contains(visited, v) {
			visited = append(visited, v)
		}
	}
	components = append(components, component)
	for i := 1; i < len(graph); i++ {
		if contains(visited, uint(i)) {
			continue
		}
		component = bfs(graph, uint(i))
		for _, v := range component {
			if !contains(visited, v) {
				visited = append(visited, v)
			}
		}
		components = append(components, component)
	}
	return components
}
