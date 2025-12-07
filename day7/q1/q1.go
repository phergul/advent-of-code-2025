package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day7/tachyon_manifold.txt")
	rows := strings.Split(strings.TrimSpace(string(data)), "\n")
	totalSplits := 0

	getPreviousBeams := func(row int) []int {
		beams := []int{}
		for i, c := range rows[row-1] {
			if c == '|' || c == 'S' {
				beams = append(beams, i)
			}
		}
		return beams
	}

	replaceRune := func(runes []rune, index int, r rune) {
		length := len(runes)
		if index < 0 || index >= length {
			return
		}
		if runes[index] == r {
			return
		}
		runes[index] = r
	}

	for i := 1; i < len(rows); i++ {
		row := rows[i]
		runes := []rune(row)

		beams := getPreviousBeams(i)
		for _, idx := range beams {
			if runes[idx] == '.' {
				runes[idx] = '|'
			}
			if runes[idx] == '^' {
				replaceRune(runes, idx+1, '|')
				replaceRune(runes, idx-1, '|')
				totalSplits++
			}
		}

		rows[i] = string(runes)
	}

	fmt.Println(totalSplits)
}
