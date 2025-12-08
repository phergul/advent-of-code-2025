package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Box struct {
	x, y, z float64
}

type Pair struct {
	b1, b2 Box
	dist   float64
}

func main() {
	data, _ := os.ReadFile("day8/coords.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	boxes := make([]Box, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x_coord, _ := strconv.ParseFloat(parts[0], 64)
		y_coord, _ := strconv.ParseFloat(parts[1], 64)
		z_coord, _ := strconv.ParseFloat(parts[2], 64)
		boxes[i] = Box{
			x: x_coord,
			y: y_coord,
			z: z_coord,
		}
	}

	distanceList := make([]Pair, 0)
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			closeness := distance(boxes[i], boxes[j])
			distanceList = append(distanceList, Pair{b1: boxes[i], b2: boxes[j], dist: closeness})
		}
	}
	sort.Slice(distanceList, func(i, j int) bool {
		return distanceList[i].dist < distanceList[j].dist
	})

	fmt.Println("Top 10 closest pairs:")
	for i := range 10 {
		fmt.Println(distanceList[i])
	}
	fmt.Println()

	circuits := make(map[int][]Box)
	idCounter := 0
	lastBoxesXCoord := [2]int{}
	for i := 0; i < len(distanceList); i++ {
		pair := distanceList[i]
		box1 := pair.b1
		box2 := pair.b2
		fmt.Println("Considering pair:", box1, box2, "with distance", pair.dist)

		if len(circuits) == 0 {
			fmt.Println("No circuits yet, creating first circuit")
			circuits[idCounter] = []Box{box1, box2}
			idCounter++
			continue
		}

		box1CircuitId, box2CircuitId := getCircuitIds(circuits, box1, box2)
		if box1CircuitId != -1 && box2CircuitId != -1 && box1CircuitId == box2CircuitId {
			fmt.Println("Both boxes already in the same circuit", box1CircuitId)
			continue
		} else if box1CircuitId != -1 && box2CircuitId != -1 && box1CircuitId != box2CircuitId {
			fmt.Println("Both boxes in different circuits", box1CircuitId, "and", box2CircuitId)
			fmt.Println("Merging circuits")
			mergedCircuit := append(circuits[box1CircuitId], circuits[box2CircuitId]...)
			delete(circuits, box1CircuitId)
			delete(circuits, box2CircuitId)
			circuits[idCounter] = mergedCircuit
			idCounter++
			continue
		} else if box1CircuitId != -1 && box2CircuitId == -1 {
			if len(circuits) == 1 {
				lastBoxesXCoord[0] = int(box1.x)
				lastBoxesXCoord[1] = int(box2.x)
			}
			fmt.Println("Box1 in circuit", box1CircuitId, "adding Box2 to it")
			circuits[box1CircuitId] = append(circuits[box1CircuitId], box2)
		} else if box1CircuitId == -1 && box2CircuitId != -1 {
			if len(circuits) == 1 {
				lastBoxesXCoord[0] = int(box1.x)
				lastBoxesXCoord[1] = int(box2.x)
			}
			fmt.Println("Box2 in circuit", box2CircuitId, "adding Box1 to it")
			circuits[box2CircuitId] = append(circuits[box2CircuitId], box1)
		} else {
			fmt.Println("Neither box in any circuit, creating new circuit")
			circuits[idCounter] = []Box{box1, box2}
			idCounter++
		}
		fmt.Println()
	}

	totalBoxes := 0
	for id, circuit := range circuits {
		fmt.Println("Circuit", id, "has boxes:", circuit)
		totalBoxes += len(circuit)
	}
	fmt.Println("Formed", len(circuits), "circuits")
	fmt.Println("Total circuits", len(circuits)+len(boxes)-totalBoxes)
	fmt.Println("Total boxes in circuits:", totalBoxes, "should be", len(boxes))

	boxesIn3Biggest := 1
	sortedCircuits := make([][]Box, 0, len(circuits))
	for _, circuit := range circuits {
		sortedCircuits = append(sortedCircuits, circuit)
	}
	sort.Slice(sortedCircuits, func(i, j int) bool {
		return len(sortedCircuits[i]) > len(sortedCircuits[j])
	})

	fmt.Println("Sizes of circuits in descending order:")
	for i, circuit := range sortedCircuits {
		fmt.Println("Circuit", i, "size:", len(circuit))
	}
	for i := 0; i < 3 && i < len(sortedCircuits); i++ {
		boxesIn3Biggest *= len(sortedCircuits[i])
	}
	fmt.Println("Total boxes in 3 biggest circuits:", boxesIn3Biggest)

	fmt.Println("Last considered boxes' X coordinates:", lastBoxesXCoord)
	fmt.Println("Product of their X coordinates:", lastBoxesXCoord[0]*lastBoxesXCoord[1])
}

func getCircuitIds(circuits map[int][]Box, box1, box2 Box) (int, int) {
	id1, id2 := -1, -1
	for id, circuit := range circuits {
		if idx := slices.Index(circuit, box1); idx != -1 {
			id1 = id
		}
		if idx := slices.Index(circuit, box2); idx != -1 {
			id2 = id
		}
	}
	return id1, id2
}

func distance(b1, b2 Box) float64 {
	return math.Sqrt((b1.x-b2.x)*(b1.x-b2.x) + (b1.y-b2.y)*(b1.y-b2.y) + (b1.z-b2.z)*(b1.z-b2.z))
}
