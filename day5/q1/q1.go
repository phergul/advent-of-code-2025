package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day5/ingredients.txt")
	parts := strings.SplitAfter(string(data), "\n\n")
	idRangeStrings := strings.Split(strings.TrimSpace(parts[0]), "\n")

	idRanges := make([][2]int, len(idRangeStrings))
	for i, r := range idRangeStrings {
		bounds := strings.Split(r, "-")
		idRanges[i][0], _ = strconv.Atoi(bounds[0])
		idRanges[i][1], _ = strconv.Atoi(bounds[1])
	}

	ids := strings.Split(strings.TrimSpace(parts[1]), "\n")

	freshCount := 0
	for _, id := range ids {
		idInt, _ := strconv.Atoi(id)
		for _, idBounds := range idRanges {
			if idInt >= idBounds[0] && idInt <= idBounds[1] {
				freshCount++
				break
			}
		}
	}
	fmt.Println(freshCount)
}
