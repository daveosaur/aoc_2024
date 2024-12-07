package main

import (
	"testing"
)

var testinput string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func TestP1(t *testing.T) {

	result := part1(testinput)
	if result != 3749 {
		t.Errorf("failed. got %d, should be 3749", result)
	}
}
func TestP2(t *testing.T) {

	result := part2(testinput)
	if result != 11387 {
		t.Errorf("failed. got %d, should be 11387", result)
	}
}
