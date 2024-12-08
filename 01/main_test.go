package main

import (
	"testing"
)

var testinput string = `3   4
4   3
2   5
1   3
3   9
3   3
`

func TestP1(t *testing.T) {

	result := part1(testinput)
	if result != 11 {
		t.Errorf("failed. got %d, should be 11", result)
	}
}
func TestP2(t *testing.T) {

	result := part2(testinput)
	if result != 31 {
		t.Errorf("failed. got %d, should be 31", result)
	}
}
