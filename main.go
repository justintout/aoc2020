package main

import "fmt"

func main() {
	day1part1()
	day1part2()
}

func day1part1() {
	// report := []int{1721, 979, 366, 299, 675, 1456}
	report, err := ReadInputFromFile("./input/day1")
	if err != nil {
		panic(err)
	}
	var answer int
	for a := 0; a < len(report); a++ {
		for b := a + 1; b < len(report); b++ {
			if report[a]+report[b] == 2020 {
				answer = report[a] * report[b]
			}
		}
	}

	if answer == 0 {
		panic("no answer found")
	}

	fmt.Println(answer)
}

func day1part2() {
	// report := []int{1721, 979, 366, 299, 675, 1456}
	report, err := ReadInputFromFile("./input/day1")
	if err != nil {
		panic(err)
	}
	var answer int
	for a := 0; a < len(report); a++ {
		for b := a + 1; b < len(report); b++ {
			for c := b + 1; c < len(report); c++ {
				if report[a]+report[b]+report[c] == 2020 {
					answer = report[a] * report[b] * report[c]
				}
			}
		}
	}

	if answer == 0 {
		panic("no answer found")
	}

	fmt.Println(answer)
}
