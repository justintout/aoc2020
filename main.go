package main

import "fmt"

func main() {
	answers, err := Day1("./input/day1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 1\n\tPart 1: %d\n\tPart 2: %d\n", answers[0], answers[1])
}
