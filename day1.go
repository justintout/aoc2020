package main

import (
	"fmt"
	"strconv"

	"gonum.org/v1/gonum/stat/combin"
)

func day1(inputs []string) (int, int, error) {
	var numbers []int
	var part1, part2 int
	for _, line := range inputs {
		n, err := strconv.Atoi(line)
		if err != nil {
			return part1, part2, fmt.Errorf("error parsing input line: %s", err.Error())
		}
		numbers = append(numbers, n)
	}
	var err error
	if part1, err = findSum(2, 2020, numbers); err != nil {
		return part1, part2, fmt.Errorf("error in part 1: %s", err.Error())
	}
	if part2, err = findSum(3, 2020, numbers); err != nil {
		return part1, part2, fmt.Errorf("error in part 2: %s", err.Error())
	}

	return part1, part2, nil
}

func findSum(n int, target int, inputs []int) (int, error) {
	var answer int
	cs := combin.Combinations(len(inputs), n)
	for _, c := range cs {
		sum := 0
		var is []int
		for _, idx := range c {
			is = append(is, inputs[idx])
		}
		for _, idx := range c {
			sum = sum + inputs[idx]
		}
		if sum == target {
			answer = inputs[c[0]]
			for i := 1; i < len(c); i++ {
				answer = answer * inputs[c[i]]
			}
			return answer, nil
		}
	}

	return answer, fmt.Errorf("no answer found")
}
