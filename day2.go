package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day2 solves the second day for AoC 2020
func Day2(path string) ([2]int, error) {
	var answers [2]int
	inputs, err := ReadStringsFromFile(path)
	if err != nil {
		return answers, err
	}

	answers[0], err = day2part1(inputs)
	if err != nil {
		return answers, err
	}

	answers[1], err = day2part2(inputs)
	if err != nil {
		return answers, err
	}

	return answers, err
}

func day2part1(inputs []string) (int, error) {
	var valid int
	for _, input := range inputs {
		s := strings.Split(input, ":")   // policy is at 0, password is at 1
		ss := strings.Split(s[0], " ")   // range is at 0, letter is at 1
		sss := strings.Split(ss[0], "-") // lower is at 0, upper is at 1
		lower, err := strconv.Atoi(sss[0])
		if err != nil {
			return valid, fmt.Errorf("error parsing lower bound in policy: %s", err.Error())
		}
		upper, err := strconv.Atoi(sss[1])
		if err != nil {
			return valid, fmt.Errorf("error parsing upper bound in policy: %s", err.Error())
		}
		var count int
		for _, c := range s[1] {
			if string(c) == ss[1] {
				count = count + 1
			}
		}
		if count >= lower && count <= upper {
			valid = valid + 1
		}
	}
	return valid, nil
}

func day2part2(inputs []string) (int, error) {
	var valid int

	for _, input := range inputs {
		s := strings.Split(input, ":")   // policy is at 0, password is at 1
		ss := strings.Split(s[0], " ")   // range is at 0, letter is at 1
		sss := strings.Split(ss[0], "-") // first is at 0, second is at 1
		first, err := strconv.Atoi(sss[0])
		if err != nil {
			return valid, fmt.Errorf("error parsing first position in policy: %s", err.Error())
		}
		second, err := strconv.Atoi(sss[1])
		if err != nil {
			return valid, fmt.Errorf("error parsing first position in policy: %s", err.Error())
		}

		if (string(s[1][first]) == ss[1] && string(s[1][second]) != ss[1]) || (string(s[1][first]) != ss[1] && string(s[1][second]) == ss[1]) {
			valid = valid + 1
		}
	}
	return valid, nil
}
