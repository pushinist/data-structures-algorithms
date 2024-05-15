package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func graphFromFile(pathToFile string) [][]uint {
	file, _ := os.Open(pathToFile)
	defer file.Close()
	vertice := 0
	adjacencyList := make([][]uint, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		adjacencyList = append(adjacencyList, make([]uint, 0))
		line := scanner.Text()
		relations := strings.Fields(line)
		for i, v := range relations {
			number, _ := strconv.Atoi(v)
			if number != 0 {
				adjacencyList[vertice] = append(adjacencyList[vertice], uint(i))
			}
		}
		vertice++
	}
	return adjacencyList
}
