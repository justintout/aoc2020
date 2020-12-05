package main

import "testing"

func TestDay5(t *testing.T) {
	inputs := []string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	expected1 := 820
	expected2 := 0

	result1, result2, err := day5(inputs)
	if err != nil {
		t.Error(err)
	}

	if result1 != expected1 {
		t.Errorf("error solving part 1, got: %d, want: %d", result1, expected1)
	}
	if result2 != expected2 {
		t.Errorf("error solving part 2, got: %d, want: %d", result2, expected2)
	}
}

func TestSeatID(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "FBFBBFFRLR",
			expected: 357,
		},
		{
			input:    "BFFFBBFRRR",
			expected: 567,
		},
		{
			input:    "FFFBBBFRRR",
			expected: 119,
		},
		{
			input:    "BBFFBBFRLL",
			expected: 820,
		},
	}

	for _, tt := range tests {
		result := seatID(tt.input)
		if result != tt.expected {
			t.Errorf("error solving seat ID %s, got: %d, want: %d", tt.input, result, tt.expected)
		}
	}
}
