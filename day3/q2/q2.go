package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var total int = 0

func main() {
	data, _ := os.ReadFile("day3/banks.txt")
	banks := strings.SplitSeq(strings.TrimSpace(string(data)), "\n")

	for bank := range banks {
		largestValues := [12][2]int{}
		largestValues[11] = largestFromIndexRange(-1, len(bank), 11, bank)
		for i := 10; i >= 0; i-- {
			largestValues[i] = largestFromIndexRange(largestValues[i+1][0], len(bank), i, bank)
		}

		bankString := ""
		for i := len(largestValues) - 1; i >= 0; i-- {
			bankString += strconv.Itoa(largestValues[i][1])
		}
		final, _ := strconv.Atoi(bankString)
		total += final
	}
	fmt.Println(total)
}

func largestFromIndexRange(prevIndex, maxIndex, counter int, input string) [2]int {
	largest := [2]int{}
	for i := prevIndex + 1; i < maxIndex-counter; i++ {
		c, _ := strconv.Atoi(string(input[i]))
		if c > largest[1] {
			largest[0] = i
			largest[1] = c
		}
	}
	return largest
}
