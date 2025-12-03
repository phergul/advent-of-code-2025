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
		largest := [2]int{}
		for i := 0; i < len(bank)-1; i++ {
			c, _ := strconv.Atoi(string(bank[i]))
			if c > largest[1] {
				largest[0] = i
				largest[1] = c
			}
		}

		largestAfterFirst := [2]int{}
		for i := largest[0] + 1; i < len(bank); i++ {
			c, _ := strconv.Atoi(string(bank[i]))
			if c > largestAfterFirst[1] {
				largestAfterFirst[0] = i
				largestAfterFirst[1] = c
			}
		}

		v, _ := strconv.Atoi(fmt.Sprintf("%d%d", largest[1], largestAfterFirst[1]))
		total += v
	}
	fmt.Println(total)
}
