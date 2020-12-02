package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var days []solver = []solver{day1, day2}

func main() {
	for i, d := range days {
		name := fmt.Sprintf("Day %d", i+1)
		path := fmt.Sprintf("./input/day%d", i+1)
		input, err := readInput(path)
		if err != nil {
			log.Fatalf("error reading %s input: %s", name, err.Error())
		}
		part1, part2, err := d(input)
		if err != nil {
			log.Fatalf("error solving %s: %s", name, err.Error())
		}
		fmt.Printf("%s\n\tPart 1: %d\n\tPart 2: %d\n", name, part1, part2)
	}
}

type solver func(inputs []string) (int, int, error)

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var strings []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		strings = append(strings, s.Text())
	}
	return strings, err
}
