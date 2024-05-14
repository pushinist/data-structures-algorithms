package main

import "slices"

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
			if !contains(visited, v) {
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

func getComps(graph [][]uint) [][]uint {
	components := make([][]uint, 0)
	var component []uint
	visited := make([]uint, 0)
	for i := range graph {
		component = bfs(graph, uint(i))
		for _, v := range component {
			if !contains(visited, v) {
				visited = append(visited, v)
			}
		}
		slices.Sort(component)
		if !containsArr(components, component) {
			components = append(components, component)
		}
	}
	return components
}
