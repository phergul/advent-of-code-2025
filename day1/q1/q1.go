package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_NUM = 100

var (
	start = 50
	count = 0
)

func main() {
	data, _ := os.ReadFile("day1/rotations.txt")

	rotations := strings.SplitSeq(strings.TrimSpace(string(data)), "\n")

	for rotation := range rotations {
		dir := rotation[:1]
		distance, _ := strconv.Atoi(rotation[1:])

		switch dir {
		case "L":
			start -= distance
			for start < 0 {
				start = MAX_NUM + start
			}
		case "R":
			start += distance
			start %= MAX_NUM
		}

		if start == 0 {
			count++
		}
	}

	fmt.Println(count)
}
