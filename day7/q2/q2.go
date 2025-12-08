package main

import (
	"bytes"
	"fmt"
	"os"
)

var cache = make(map[string]int)
var grid [][]byte
var numRows, numCols int

func main() {
	data, _ := os.ReadFile("day7/tachyon_manifold.txt")
	lines := bytes.Split(bytes.TrimSpace(data), []byte("\n"))

	numRows = len(lines)
	numCols = len(lines[0])
	grid = make([][]byte, numRows)
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		copy(grid[i], line)
	}

	start := -1
	for c, ch := range grid[0] {
		if ch == 'S' {
			start = c
			break
		}
	}

	totalTimelines := countTimelines(1, start)

	fmt.Println(totalTimelines)
}

func countTimelines(row, col int) int {
	if row >= numRows {
		return 1
	}

	if col < 0 || col >= numCols {
		return 0
	}

	cacheKey := fmt.Sprintf("%d,%d", row, col)
	if v, ok := cache[cacheKey]; ok {
		return v
	}

	count := 0
	c := grid[row][col]

	if c == '^' {
		count = countTimelines(row+1, col-1) + countTimelines(row+1, col+1)
	} else {
		count = countTimelines(row+1, col)
	}

	cache[cacheKey] = count
	return count
}
