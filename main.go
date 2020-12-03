package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var cookie *string = flag.String("cookie", "", "adventofcode.com session cookie. if supplied, the script will fetch inputs from adventofcode.com")
var day *int = flag.Int("day", 0, "solve only the given day, instead of all unlocked days")

var days []solver = []solver{day1, day2, day3}

func main() {
	flag.Parse()

	if *day != 0 && *day <= len(days) {
		solve(*day, days[*day-1])
		return
	}
	for i, d := range days {
		solve(i+1, d)
	}
}

type solver func(inputs []string) (int, int, error)

func solve(day int, d solver) {
	name := fmt.Sprintf("Day %d", day)
	var input []string
	var err error
	if *cookie != "" {
		if input, err = fetchInput(day, *cookie); err != nil {
			log.Fatalf("error fetching input for %s: %s", name, err.Error())
		}
	} else {
		path := fmt.Sprintf("./input/day%d", day)
		if input, err = readInput(path); err != nil {
			log.Fatalf("error reading %s input: %s", name, err.Error())
		}
	}
	part1, part2, err := d(input)
	if err != nil {
		log.Fatalf("error solving %s: %s", name, err.Error())
	}
	fmt.Printf("%s\n\tPart 1: %d\n\tPart 2: %d\n", name, part1, part2)
}

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

func fetchInput(day int, cookie string) ([]string, error) {
	url := fmt.Sprintf("https://adventofcode.com/2020/day/%d/input", day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// the last line is always a blank, and will get included in our inputs
	// be sure to slice the last line
	lines := strings.Split(string(body), "\n")

	return lines[:len(lines)-1], nil
}
