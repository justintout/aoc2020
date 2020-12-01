package main

import "fmt"

// Day1 returns the solutions to day 1's puzzle, given the path to the input file
func Day1(path string) ([2]int, error) {
	input, err := ReadInputFromFile(path)
	var answers [2]int
	if err != nil {
		return answers, err
	}
	p1, err := part1(input)
	if err != nil {
		return answers, fmt.Errorf("error in part 1: %s", err.Error())
	}
	p2, err := part2(input)
	if err != nil {
		return answers, fmt.Errorf("error in part 2: %s", err.Error())
	}
	return [2]int{p1, p2}, nil
}

func part1(input []int) (int, error) {
	var answer int
	for a := 0; a < len(input); a++ {
		for b := a + 1; b < len(input); b++ {
			if input[a]+input[b] == 2020 {
				answer = input[a] * input[b]
			}
		}
	}

	var err error
	if answer == 0 {
		err = fmt.Errorf("no answer found")
	}

	return answer, err
}

func part2(input []int) (int, error) {
	var answer int
	for a := 0; a < len(input); a++ {
		for b := a + 1; b < len(input); b++ {
			for c := b + 1; c < len(input); c++ {
				if input[a]+input[b]+input[c] == 2020 {
					answer = input[a] * input[b] * input[c]
				}
			}
		}
	}

	var err error
	if answer == 0 {
		err = fmt.Errorf("no answer found")
	}

	return answer, err
}
