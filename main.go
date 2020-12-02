package main

import "fmt"

func main() {
	day1, err := Day1("./input/day1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 1\n\tPart 1: %d\n\tPart 2: %d\n", day1[0], day1[1])

	day2, err := Day2("./input/day2")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day2\n\tPart 1: %d\n\tPart 2: %d\n", day2[0], day2[1])
}

type Solver struct {
	path string
	
}