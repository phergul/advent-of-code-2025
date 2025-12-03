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
			if idLen%2 == 1 {
				start++
				continue
			}

			half1 := idString[:idLen/2]
			half2 := idString[idLen/2:]
			if half1 == half2 {
				count += start
			}

			start++
		}
	}
	fmt.Println(count)
}
