package main

import "testing"

func TestDay2(t *testing.T) {
	inputs := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}

	expected1 := 2
	expected2 := 1

	result1, result2, err := day2(inputs)
	if err != nil {
		t.Error(err)
	}
	if expected1 != result1 {
		t.Errorf("error in part 1, got: %d, want: %d", result1, expected1)
	}

	if expected2 != result2 {
		t.Errorf("error in part 2, got: %d, want: %d", result2, expected2)
	}
}
