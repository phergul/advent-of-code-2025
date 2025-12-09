package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

func main() {
	data, _ := os.ReadFile("day9/tiles.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	coords := make([]Coord, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x_coord, _ := strconv.Atoi(parts[0])
		y_coord, _ := strconv.Atoi(parts[1])
		coords[i] = Coord{x: x_coord, y: y_coord}
	}

	calcMap := make(map[string]int)
	for i := range coords {
		for j := range coords {
			if i == j {
				continue
			}
			if _, exists := calcMap[fmt.Sprintf("%d,%d-%d,%d", coords[j].x, coords[j].y, coords[i].x, coords[i].y)]; exists {
				continue
			}
			calcMap[fmt.Sprintf("%d,%d-%d,%d", coords[i].x, coords[i].y, coords[j].x, coords[j].y)] = coords[i].area(coords[j])
		}
	}

	largestSquare := 0
	for _, v := range calcMap {
		if v > largestSquare {
			largestSquare = v
		}
	}
	fmt.Println(largestSquare)
}

func (c Coord) area(other Coord) int {
	floatx1 := float64(c.x)
	floaty1 := float64(c.y)
	floatx2 := float64(other.x)
	floaty2 := float64(other.y)
	if floatx2 < floatx1 {
		floatx1++
	} else {
		floatx2++
	}
	if floaty2 < floaty1 {
		floaty1++
	} else {
		floaty2++
	}
	return int(math.Abs(floatx2-floatx1) * math.Abs(floaty2-floaty1))
}
