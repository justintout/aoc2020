package main

import "testing"

func TestDay1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		input := []int{1721, 979, 366, 299, 675, 1456}
		expected := 514579
		result, err := day1part1(input)
		if err != nil {
			t.Error(err)
		}

		if expected != result {
			t.Errorf("incorrect result; got: %d; want: %d", result, expected)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		input := []int{1721, 979, 366, 299, 675, 1456}
		expected := 241861950
		result, err := day1part2(input)
		if err != nil {
			t.Error(err)
		}

		if expected != result {
			t.Errorf("incorrect result; got: %d; want: %d", result, expected)
		}
	})

	t.Run("Day 1", func(t *testing.T) {
		expected := [2]int{514579, 241861950}
		answers, err := Day1("./input/day1-example")
		if err != nil {
			t.Error(err)
		}

		for i := 0; i < 2; i++ {
			if expected[i] != answers[i] {
				t.Errorf("wrong answer: %d; got: %d; want: %d", i, answers[i], expected[i])
			}
		}
	})
}
