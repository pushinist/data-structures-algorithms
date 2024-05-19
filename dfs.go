package main

func dfs(graph [][]uint, start uint) []uint {
	visited := make([]uint, 0)
	stack := Stack{}
	stack.Push(start)
	for len(stack) > 0 {
		visited = append(visited, stack.Pop())
		for _, v := range graph[visited[len(visited)-1]] {
			if !contains(visited, v) && !contains(stack, v) {
				stack.Push(v)
			}
		}
	}
	return visited
}

func dfsGetComps(graph [][]uint) [][]uint {
	component := dfs(graph, 0)
	components := make([][]uint, 0)
	visited := make([]uint, 0)
	for _, v := range component {
		if !contains(visited, v) {
			visited = append(visited, v)
		}
	}
	components = append(components, component)
Search:
	for i := 1; i < len(graph); i++ {
		if contains(visited, uint(i)) {
			continue
		}
		component = dfs(graph, uint(i))
		for _, value := range component {
			if len(graph[value]) == 0 {
				continue Search
			}
		}
		for _, v := range component {
			if !contains(visited, v) {
				visited = append(visited, v)
			}
		}
		components = append(components, component)
	}
	return components
}
