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

func main() {
	data, _ := os.ReadFile("day4/rolls.txt")
	rows := strings.Split(strings.TrimSpace(string(data)), "\n")
	colLen := len(rows)
	rowLen := len(rows[0])

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

	// for i, v := range grid {
	// 	if v {
	// 		fmt.Print("@")
	// 	} else {
	// 		fmt.Print(".")
	// 	}
	// 	if (i+1)%rowLen == 0 {
	// 		fmt.Println()
	// 	}
	// }
	// fmt.Println()

	inaccessibleCount := 0
	for i := range grid {
		val := grid[i]
		if !val {
			inaccessibleCount++
			continue
		}
		adjacentGrid := make([]bool, 8)

		// fmt.Printf("Row: %d, Col: %d\n", i/rowLen+1, i%rowLen+1)
		if i > rowLen {
			adjacentGrid[Top] = grid[i-rowLen]
		}
		if i <= (len(grid)-1)-rowLen {
			adjacentGrid[Bottom] = grid[i+rowLen]
		}
		// if i%(rowLen) != rowLen && i < len(grid)-1 {
		if i < len(grid)-1 && i%(rowLen) != rowLen-1 {
			adjacentGrid[Right] = grid[i+1]
		}
		if i > 0 && i%(rowLen) != 0 {
			adjacentGrid[Left] = grid[i-1]
		}
		if i > rowLen+1 && i%(rowLen) != 0 {
			adjacentGrid[TopLeft] = grid[i-rowLen-1]
		}
		// if i > rowLen && i%(rowLen) != rowLen {
		if i > rowLen-1 && i%(rowLen) != rowLen-1 {
			adjacentGrid[TopRight] = grid[i-rowLen+1]
		}
		// if i <= (len(grid)-1)-(rowLen-1) && i%(rowLen) != 0 {
		if i <= (len(grid)-1)-(rowLen-1) && i%(rowLen) != 0 {
			adjacentGrid[BottomLeft] = grid[i+rowLen-1]
		}
		// if i <= (len(grid)-1)-(rowLen+1) && i%(rowLen) != rowLen {
		if i <= (len(grid)-1)-(rowLen+1) && i%(rowLen) != rowLen-1 {
			adjacentGrid[BottomRight] = grid[i+rowLen+1]
		}

		// for j, v := range adjacentGrid {
		// 	if v {
		// 		fmt.Print("@")
		// 	} else {
		// 		fmt.Print(".")
		// 	}
		// 	if j%3 == 2 && j < 4 {
		// 		fmt.Println()
		// 	}
		// 	if j == 3 {
		// 		fmt.Print(" ")
		// 	}
		// 	if j == 4 {
		// 		fmt.Print("\n")
		// 	}
		// }
		// fmt.Println()

		trueCount := 0
		for _, v := range adjacentGrid {
			if v {
				trueCount++
			}
			if trueCount >= 4 {
				inaccessibleCount++
				break
			}
		}
	}
	fmt.Println(len(grid) - inaccessibleCount)
}
