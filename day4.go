package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var validators = map[string]func(input string) bool{
	"byr": func(input string) bool { return validateRange(input, 1920, 2002) },
	"iyr": func(input string) bool { return validateRange(input, 2010, 2020) },
	"eyr": func(input string) bool { return validateRange(input, 2020, 2030) },
	"hgt": func(input string) bool { return validateHeight(input) },
	"hcl": func(input string) bool { return validateHexColor(input) },
	"ecl": func(input string) bool { return validateEyeColor(input) },
	"pid": func(input string) bool { return validatePID(input) },
	"cid": func(input string) bool { return true },
}

func day4(inputs []string) (int, int, error) {
	var part1, part2 int

	var passports []string
	var passport string
	for i, input := range inputs {
		if i == len(inputs)-1 && input != "" {
			passport = strings.Join([]string{passport, input}, " ")
			input = ""
		}
		if input == "" {
			passports = append(passports, passport)
			passport = ""
			continue
		}
		passport = strings.Join([]string{passport, input}, " ")
	}
PASSPORT:
	for _, passport := range passports {
		for _, r := range required {
			if !strings.Contains(passport, r) {
				continue PASSPORT
			}
		}
		part1 = part1 + 1
		if valid(passport) {
			fmt.Printf("valid:   %v\n", passport)
			part2 = part2 + 1
		}
	}

	return part1, part2, nil
}

func valid(passport string) bool {
	fields := strings.Split(passport, " ")
	for _, f := range fields {
		if f == "" {
			continue
		}
		ff := strings.Split(f, ":")
		if !validators[ff[0]](ff[1]) {
			fmt.Printf("invalid: %s (%s:%s)\n", passport, ff[0], ff[1])
			return false
		}
	}
	return true
}

func validateRange(number string, min, max int) bool {
	n, err := strconv.Atoi(number)
	if err != nil {
		return false
	}
	return min <= n && n <= max
}

func validateHeight(height string) bool {
	switch {
	case strings.Contains(height, "cm"):
		return validateRange(height[:len(height)-2], 150, 193)
	case strings.Contains(height, "in"):
		return validateRange(height[:len(height)-2], 59, 76)
	}
	return false
}

func validateHexColor(color string) bool {
	matched, err := regexp.MatchString(`#[0-9a-f]{6}`, color)
	return matched && err == nil
}

var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func validateEyeColor(color string) bool {
	for _, c := range eyeColors {
		if color == c {
			return true
		}
	}
	return false
}

func validatePID(pid string) bool {
	matched, err := regexp.MatchString(`[0-9]{9}`, pid)
	return matched && err == nil
}
