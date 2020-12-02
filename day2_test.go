package main

import "testing"

func TestDay2(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		input := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
		expected := 2

		result, err := day2part1(input)
		if err != nil {
			t.Error(err)
		}

		if expected != result {
			t.Errorf("incorrect answer; got: %d; want: %d", result, expected)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		input := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
		expected := 1

		result, err := day2part2(input)
		if err != nil {
			t.Error(err)
		}

		if expected != result {
			t.Errorf("incorrect answer; got: %d; want: %d", result, expected)
		}
	})

}
