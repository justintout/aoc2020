package main

import "testing"

func TestReadInputFromFile(t *testing.T) {
	path := "./input/day1-example"
	expected := []int{1721, 979, 366, 299, 675, 1456}
	result, err := ReadInputFromFile(path)
	if err != nil {
		t.Errorf("failed reading fixture file: %s", err.Error())
	}

	if len(expected) != len(result) {
		t.Errorf("unequal lengths; got: %d; want %d", len(result), len(expected))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] {
			t.Errorf("unequal element %d; got: %d; want: %d", i, result[i], expected[i])
		}
	}
}
