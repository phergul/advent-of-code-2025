package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Bound struct {
	start, end uint64
}

type Bounds []Bound

var freshCount uint64

func main() {
	data, _ := os.ReadFile("day5/ingredients.txt")

	parts := strings.SplitAfter(string(data), "\n\n")
	idRangeStrings := strings.Split(strings.TrimSpace(parts[0]), "\n")

	idRanges := make([][2]uint64, len(idRangeStrings))
	for i, r := range idRangeStrings {
		bounds := strings.Split(r, "-")
		idRanges[i][0], _ = strconv.ParseUint(bounds[0], 0, 64)
		idRanges[i][1], _ = strconv.ParseUint(bounds[1], 0, 64)
	}

	normalisedIdRanges := make(Bounds, 0, len(idRanges))
	for _, idBounds := range idRanges {
		normalisedRange := Bound{start: idBounds[0], end: idBounds[1]}
		fmt.Println("Checking:", idBounds)
		newBound, completeOverlap := normalisedIdRanges.checkOverlap(idBounds[0], idBounds[1])
		if completeOverlap {
			fmt.Print("Discarded range, fully contained\n\n")
			continue
		}
		if newBound.start != 0 {
			normalisedRange.start = newBound.start
		}
		if newBound.end != 0 {
			normalisedRange.end = newBound.end
		}
		fmt.Print("Normalised to:", normalisedRange)
		if normalisedRange.start > normalisedRange.end {
			fmt.Printf("\nDiscarded range, diff is %d\n\n", normalisedRange.end-normalisedRange.start)
			continue
		}
		fmt.Print("\n\n")

		freshCount += normalisedRange.end - normalisedRange.start + 1
		normalisedIdRanges = append(normalisedIdRanges, normalisedRange)
	}

	normalisedIdRanges.sanityCheck()
	fmt.Println(freshCount)
}

func (b *Bounds) checkOverlap(start, end uint64) (Bound, bool) {
	newBound := Bound{}
	for _, bound := range *b {
		if bound.start > start && end > bound.end {
			fmt.Printf("bound %v is fully contained in the new bound {%d-%d}. removing old bound...\n", bound, start, end)
			freshCount -= bound.end - bound.start + 1
			*b = slices.Delete(*b, slices.Index(*b, bound), slices.Index(*b, bound)+1)
		}
		if (bound.start <= start) && (end <= bound.end) {
			fmt.Println("Fully Overlaps with existing bound:", bound)
			return newBound, true
		}
		if bound.start <= start && start <= bound.end {
			fmt.Println("Start Overlaps with existing bound:", bound)
			newBound.start = bound.end + 1
		}
		if bound.start <= end && end <= bound.end {
			fmt.Println("End Overlaps with existing bound:", bound)
			newBound.end = bound.start - 1
		}
	}
	return newBound, false
}

func (b Bounds) sanityCheck() {
	fmt.Print("Sanity checking bounds...")
	bCopy := make(Bounds, len(b))
	copy(bCopy, b)
	sort.Slice(bCopy, func(i, j int) bool {
		return bCopy[i].start < bCopy[j].start
	})
	for i, bound := range bCopy {
		if i == len(bCopy)-1 {
			break
		}
		nextBound := bCopy[i+1]
		if bound.end >= nextBound.start {
			//how did we get here...
			fmt.Println("Bounds overlap detected:", bound, nextBound)
		}
	}
	fmt.Print("Done\n\n")
}
