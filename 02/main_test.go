package main

import (
	"testing"
)

var testinput string = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestP1(t *testing.T) {

	result := part1(testinput)
	if result != 2 {
		t.Errorf("failed. got %d, should be 2", result)
	}
}
func TestP2(t *testing.T) {

	result := part2(testinput)
	if result != 4 {
		t.Errorf("failed. got %d, should be 4", result)
	}
}
