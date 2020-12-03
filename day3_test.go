package main

import "testing"

func TestDay3(t *testing.T) {
	input := []string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#", ".#...##..#.", "..#.##.....", ".#.#.#....#", ".#........#", "#.##...#...", "#...##....#", ".#..#...#.#"}

	expected1 := 7
	expected2 := 336

	result1, result2, err := day3(input)
	if err != nil {
		t.Error(err)
	}

	if expected1 != result1 {
		t.Errorf("error solving part 1, got: %d, want: %d", result1, expected1)
	}

	if expected2 != result2 {
		t.Errorf("error solving part 2, got: %d, want: %d", result2, expected2)
	}
}
