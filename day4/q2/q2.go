package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	TopLeft int = iota
	Top
	TopRight
	Left
	Right
	BottomLeft
	Bottom
	BottomRight
)

type Grid struct {
	flattened      []bool
	rowLen, colLen int
}

type GridCell struct {
	index      int
	accessible bool
}

func main() {
	data, _ := os.ReadFile("day4/rolls.txt")
	rows := strings.Split(strings.TrimSpace(string(data)), "\n")
	colLen := len(rows)
	rowLen := len(rows[0])

	removed := 0
	accessible := 1

	for accessible > 0 {
		grid := makeGrid(rows, colLen, rowLen)
		cells := make([]GridCell, colLen*rowLen)
		accessible, cells = checkGridAccessability(&grid)

		for _, c := range cells {
			if c.accessible && grid.flattened[c.index] {
				removed++
				removeRoll(&rows, c.index, grid.rowLen)
			}
		}
	}
	fmt.Println(removed)
}

func makeGrid(rows []string, colLen, rowLen int) Grid {
	grid := make([]bool, colLen*rowLen)

	for j, row := range rows {
		for i := 0; i < len(row); i++ {
			c := string(row[i])
			index := (j * colLen) + i
			if c == "." {
				grid[index] = false
			} else {
				grid[index] = true
			}
		}
	}
	return Grid{
		flattened: grid,
		colLen:    colLen,
		rowLen:    rowLen,
	}
}

func checkGridAccessability(grid *Grid) (int, []GridCell) {
	cells := make([]GridCell, len(grid.flattened))
	for i := range grid.flattened {
		cell := GridCell{
			index:      i,
			accessible: true,
		}
		val := grid.flattened[i]
		if !val {
			continue
		}
		adjacentGrid := make([]bool, 8)

		if i > grid.rowLen {
			adjacentGrid[Top] = grid.flattened[i-grid.rowLen]
		}
		if i <= (len(grid.flattened)-1)-grid.rowLen {
			adjacentGrid[Bottom] = grid.flattened[i+grid.rowLen]
		}
		if i < len(grid.flattened)-1 && i%(grid.rowLen) != grid.rowLen-1 {
			adjacentGrid[Right] = grid.flattened[i+1]
		}
		if i > 0 && i%(grid.rowLen) != 0 {
			adjacentGrid[Left] = grid.flattened[i-1]
		}
		if i > grid.rowLen+1 && i%(grid.rowLen) != 0 {
			adjacentGrid[TopLeft] = grid.flattened[i-grid.rowLen-1]
		}
		if i > grid.rowLen-1 && i%(grid.rowLen) != grid.rowLen-1 {
			adjacentGrid[TopRight] = grid.flattened[i-grid.rowLen+1]
		}
		if i <= (len(grid.flattened)-1)-(grid.rowLen-1) && i%(grid.rowLen) != 0 {
			adjacentGrid[BottomLeft] = grid.flattened[i+grid.rowLen-1]
		}
		if i <= (len(grid.flattened)-1)-(grid.rowLen+1) && i%(grid.rowLen) != grid.rowLen-1 {
			adjacentGrid[BottomRight] = grid.flattened[i+grid.rowLen+1]
		}

		trueCount := 0
		for _, v := range adjacentGrid {
			if v {
				trueCount++
			}
			if trueCount >= 4 {
				cell.accessible = false
				break
			}
		}

		cells[i] = cell
	}
	accessibleCount := 0
	for _, c := range cells {
		if c.accessible {
			accessibleCount++
		}
	}
	return accessibleCount, cells
}

func removeRoll(rows *[]string, index, rowLen int) {
	rowString := (*rows)[index/rowLen]
	runeSlice := []rune(rowString)
	runeSlice[index%rowLen] = '.'
	(*rows)[index/rowLen] = string(runeSlice)
}
