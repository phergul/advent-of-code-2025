package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	operator string
	values   []int
}

type problems []problem

func main() {
	data, _ := os.ReadFile("day6/problems.txt")
	rows := strings.Split(strings.TrimSpace(string(data)), "\n")
	numCols := len(strings.Split(rows[0], " "))

	problems := make(problems, numCols)
	for i, r := range rows {
		values := strings.Fields(r)
		for j, v := range values {
			if i == len(rows)-1 {
				problems[j].operator = v
				continue
			}
			value, _ := strconv.Atoi(v)
			problems[j].values = append(problems[j].values, value)
		}
	}

	total := problems.calculateTotal()
	fmt.Println(total)
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
