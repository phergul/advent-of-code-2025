package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day2/ids.txt")
	ranges := strings.Split(string(data), ",")

	count := 0
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

		for start <= end {
			idString := strconv.Itoa(start)
			idLen := len(idString)

			for i := range idString {
				pattern := idString[:i]

				repeats := float64(idLen) / float64(len(pattern))
				if repeats == float64(int64(repeats)) {
					if strings.Repeat(pattern, int(repeats)) == idString {
						count += start
						break
					}
				}
			}

			start++
		}
	}
	fmt.Println(count)
}

