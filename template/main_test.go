package main

import (
	"testing"
)

var testinput string = ``

func TestP1(t *testing.T) {

	result := part1(testinput)
	if result != 0 {
		t.Errorf("failed. got %d, should be 0", result)
	}
}
func TestP2(t *testing.T) {

	result := part2(testinput)
	if result != 0 {
		t.Errorf("failed. got %d, should be 0", result)
	}
}
