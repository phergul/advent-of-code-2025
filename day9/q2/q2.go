package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type BoundingBox struct {
	// points               []Coord
	segments []Segment
}

type Coord struct {
	x, y int
}

type Segment struct {
	start, end Coord
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

	boundingBox := createBoundingBox(coords)
	fmt.Println("Bounding box created", boundingBox)

	isRectValid := func(vertices [4]Coord, boundingBox BoundingBox) bool {
		for _, v := range vertices {
			if !isPointInside(boundingBox.segments, v) {
				return false
			}
		}

		rectSegments := []Segment{
			{vertices[0], vertices[2]},
			{vertices[2], vertices[1]},
			{vertices[1], vertices[3]},
			{vertices[3], vertices[0]},
		}

		for _, rectSeg := range rectSegments {
			for _, boundingSeg := range boundingBox.segments {
				if segmentsIntersect(rectSeg, boundingSeg) {
					return false
				}
			}
		}
		return true
	}

	validSquares := make([]int, 0)
	for key, area := range calcMap {
		vertices := getVertices(key)
		if !isRectValid(vertices, boundingBox) {
			continue
		}

		validSquares = append(validSquares, area)
	}

	largestSquare := 0
	for _, v := range validSquares {
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

func createBoundingBox(coords []Coord) BoundingBox {
	segments := []Segment{}
	for i := range coords {
		start := coords[i]
		end := coords[(i+1)%len(coords)]
		segments = append(segments, Segment{start: start, end: end})
	}
	return BoundingBox{
		segments: segments,
	}
}

//	func createBoundingBox(coords []Coord) BoundingBox {
//		formedBox := []Coord{}
//		for i := 0; i < len(coords)-1 || i+1%len(coords) == 0; i++ {
//			if coords[i].x == coords[(i+1)%len(coords)].x {
//				addedLine := addVerticalLine(coords[i], coords[(i+1)])
//				formedBox = append(formedBox, addedLine...)
//			} else if coords[i].y == coords[(i+1)%len(coords)].y {
//				addedLine := addHorizontalLine(coords[i], coords[(i+1)])
//				formedBox = append(formedBox, addedLine...)
//			}
//		}
//		//last line
//		lastCoord := coords[len(coords)-1]
//		firstCoord := coords[0]
//		if lastCoord.x == firstCoord.x {
//			addedLine := addVerticalLine(lastCoord, firstCoord)
//			formedBox = append(formedBox, addedLine...)
//		} else if lastCoord.y == firstCoord.y {
//			addedLine := addHorizontalLine(lastCoord, firstCoord)
//			formedBox = append(formedBox, addedLine...)
//		}
//
//		maxX, maxY := math.MinInt, math.MinInt
//		minX, minY := math.MaxInt, math.MaxInt
//		for _, c := range coords {
//			if c.x > maxX {
//				maxX = c.x
//			} else if c.x < minX {
//				minX = c.x
//			}
//			if c.y > maxY {
//				maxY = c.y
//			} else if c.y < minY {
//				minY = c.y
//			}
//		}
//
//		return BoundingBox{
//			points:    formedBox,
//			edges:     []int{minX, minY, maxX, maxY},
//			maxLength: maxX - minX,
//			maxHeight: maxY - minY,
//		}
//	}
//
//	func addHorizontalLine(start, end Coord) []Coord {
//		formedLine := []Coord{}
//		for i := start; i != end; {
//			if start.x < end.x {
//				for j := start.x; j <= end.x; j++ {
//					formedLine = append(formedLine, Coord{x: j, y: start.y})
//					if j == end.x-1 {
//						i = end
//					}
//				}
//			} else if start.x > end.x {
//				for j := end.x; j <= start.x; j++ {
//					formedLine = append(formedLine, Coord{x: j, y: start.y})
//					if j == start.x-1 {
//						i = end
//					}
//				}
//			}
//		}
//		return formedLine
//	}
//
//	func addVerticalLine(start, end Coord) []Coord {
//		formedLine := []Coord{}
//		for i := start; i != end; {
//			if start.y < end.y {
//				for j := start.y; j <= end.y; j++ {
//					formedLine = append(formedLine, Coord{x: start.x, y: j})
//					if j == end.y-1 {
//						i = end
//					}
//				}
//			} else if start.y > end.y {
//				for j := end.y; j <= start.y; j++ {
//					formedLine = append(formedLine, Coord{x: start.x, y: j})
//					if j == start.y-1 {
//						i = end
//					}
//				}
//			}
//		}
//		return formedLine
//	}
// func (b BoundingBox) isInBoundingBox(coord Coord) bool {
// 	if slices.Contains(b.points, coord) {
// 		return true
// 	}
//
// 	boundCheck := [4]bool{}
// 	checkMutex := sync.Mutex{}
// 	var wg sync.WaitGroup
// 	for d := -1; d <= 1; d += 2 {
// 		switch d {
// 		case -1:
// 			wg.Go(func() {
// 				for x := coord.x; x > coord.x-b.maxLength; x-- {
// 					if x < b.edges[0] {
// 						break
// 					}
// 					if slices.Contains(b.points, Coord{x: x, y: coord.y}) {
// 						checkMutex.Lock()
// 						boundCheck[0] = true
// 						checkMutex.Unlock()
// 					}
// 				}
// 			})
// 			wg.Go(func() {
// 				for y := coord.y; y > coord.y-b.maxHeight; y-- {
// 					if y < b.edges[1] {
// 						break
// 					}
// 					if slices.Contains(b.points, Coord{x: coord.x, y: y}) {
// 						checkMutex.Lock()
// 						boundCheck[1] = true
// 						checkMutex.Unlock()
// 					}
// 				}
// 			})
// 		case 1:
// 			wg.Go(func() {
// 				for x := coord.x; x < coord.x+b.maxLength; x++ {
// 					if x > b.edges[2] {
// 						break
// 					}
// 					if slices.Contains(b.points, Coord{x: x, y: coord.y}) {
// 						checkMutex.Lock()
// 						boundCheck[2] = true
// 						checkMutex.Unlock()
// 					}
// 				}
// 			})
// 			wg.Go(func() {
// 				for y := coord.y; y < coord.y+b.maxHeight; y++ {
// 					if y > b.edges[3] {
// 						break
// 					}
// 					if slices.Contains(b.points, Coord{x: coord.x, y: y}) {
// 						checkMutex.Lock()
// 						boundCheck[3] = true
// 						checkMutex.Unlock()
// 					}
// 				}
// 			})
// 		}
// 	}
// 	wg.Wait()
//
// 	return boundCheck[0] && boundCheck[1] && boundCheck[2] && boundCheck[3]
// }

func getVertices(key string) [4]Coord {
	var vertices [4]Coord
	keyParts := strings.Split(key, "-")
	for i, part := range keyParts {
		coords := strings.Split(part, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		vertices[i] = Coord{x: x, y: y}
	}
	vertices[2] = Coord{x: vertices[0].x, y: vertices[1].y}
	vertices[3] = Coord{x: vertices[1].x, y: vertices[0].y}

	return vertices
}

func isPointInside(segments []Segment, point Coord) bool {
	for _, s := range segments {
		if s.start.x == s.end.x && point.x == s.start.x {
			minY := min(s.start.y, s.end.y)
			maxY := max(s.start.y, s.end.y)
			if point.y >= minY && point.y <= maxY {
				return true
			}
		}
		if s.start.y == s.end.y && point.y == s.start.y {
			minX := min(s.start.x, s.end.x)
			maxX := max(s.start.x, s.end.x)
			if point.x >= minX && point.x <= maxX {
				return true
			}
		}
	}

	intersections := 0
	for _, s := range segments {
		if (point.y > s.start.y) != (point.y > s.end.y) {
			intersectX := (s.end.x-s.start.x)*(point.y-s.start.y)/(s.end.y-s.start.y) + s.start.x
			if point.x < intersectX {
				intersections++
			}
		}
	}
	fmt.Println("Point", point, "has", intersections, "intersections")
	return intersections%2 == 1
}

func segmentsIntersect(s1, s2 Segment) bool {
	ccw := func(p1, p2, p3 Coord) int {
		val := (p2.y-p1.y)*(p3.x-p2.x) - (p2.x-p1.x)*(p3.y-p2.y)
		if val > 0 {
			return 1
		}
		if val < 0 {
			return -1
		}
		return 0
	}

	d1 := ccw(s1.start, s1.end, s2.start)
	d2 := ccw(s1.start, s1.end, s2.end)
	d3 := ccw(s2.start, s2.end, s1.start)
	d4 := ccw(s2.start, s2.end, s1.end)

	if ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) && ((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0)) {
		return true
	}
	return false
}
