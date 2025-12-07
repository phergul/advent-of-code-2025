package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	operator  string
	rawValues []string
	values    []int
}

type problems []problem

func main() {
	data, _ := os.ReadFile("day6/problems.txt")
	rows := strings.Split(strings.TrimSpace(string(data)), "\n")
	numCols := len(strings.Fields(rows[0]))
	rowLen := len(rows[0])

	problems := make(problems, numCols)
	spaceIndexes := []int{0}
	operatorRow := rows[len(rows)-1]
	opCount := 0
	for i, op := range operatorRow {
		if i < 1 {
			problems[opCount].operator = string(op)
			opCount++
			continue
		}
		if op != ' ' {
			spaceIndexes = append(spaceIndexes, i-1)
			problems[opCount].operator = string(op)
			opCount++
		}
	}
	spaceIndexes = append(spaceIndexes, rowLen)
	for i := 0; i < len(rows)-1; i++ {
		problemValues := []string{}
		for j := 1; j < len(spaceIndexes); j++ {
			if j == 1 {
				problemValues = append(problemValues, rows[i][spaceIndexes[j-1]:spaceIndexes[j]])
				continue
			}
			problemValues = append(problemValues, rows[i][spaceIndexes[j-1]+1:spaceIndexes[j]])
		}
		for k, v := range problemValues {
			v = strings.ReplaceAll(v, " ", "-")
			problems[k].rawValues = append(problems[k].rawValues, v)
		}
	}

	problems.parseValues()
	total := problems.calculateTotal()
	fmt.Println(total)
}

func (p problems) parseValues() {
	for i := range p {
		numValues := len(p[i].rawValues)
		valLen := len(p[i].rawValues[0])
		for charIdx := range valLen {
			intermediateNums := make([]string, valLen)
			for vIdx := range numValues {
				intermediateNums[charIdx] += string(p[i].rawValues[vIdx][charIdx])
			}

			for idx, strNum := range intermediateNums {
				intermediateNums[idx] = strings.Trim(strNum, "-")
			}
			numStr := strings.Join(intermediateNums, "")
			num, _ := strconv.Atoi(numStr)
			p[i].values = append(p[i].values, num)
		}
	}
}

func (p problems) calculateTotal() int {
	t := 0
	for i := range p {
		switch p[i].operator {
		case "+":
			t += sumSliceValues(p[i].values)
		case "*":
			t += multiplySliceValues(p[i].values)
		}
	}
	return t
}

func sumSliceValues(values []int) int {
	t := 0
	for _, v := range values {
		t += v
	}
	return t
}

func multiplySliceValues(values []int) int {
	t := 1
	for _, v := range values {
		t *= v
	}
	return t
}
