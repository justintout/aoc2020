package main

import (
	"bufio"
	"os"
	"strconv"
)

// ReadNumbersFromFile reads the typical AoC input format as
// a text file, returning an array of the numbers inside.
func ReadNumbersFromFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	s := bufio.NewScanner(file)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}

	return numbers, err
}

// ReadStringsFromFile reads the typical AoC input format
// as a text file, returning an array of strings (each line)
func ReadStringsFromFile(path string) ([]string, error) {
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
