package main

func contains(s []uint, e uint) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsArr(s [][]uint, e []uint) bool {
	for _, a := range s {
		if compare(a, e) {
			return true
		}
	}
	return false
}

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

func compare(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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
		component = dfs(graph, uint(i))
		for _, v := range component {
			if !contains(visited, v) {
				visited = append(visited, v)
			}
		}
		components = append(components, component)
	}
	return components

}

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
