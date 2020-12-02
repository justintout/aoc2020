package main

import "testing"

func TestDay1(t *testing.T) {
	inputs := []string{"1721", "979", "366", "299", "675", "1456"}

	expected1 := 514579
	expected2 := 241861950

	result1, result2, err := day1(inputs)
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
