package main

import (
	"testing"
)

var testinput string = ``

func TestP1(t *testing.T) {

	result := part1(testinput)
	target := 0
	if result != target {
		t.Errorf("failed. got %d, should be %d\n", result, target)
	}
}
func TestP2(t *testing.T) {

	result := part2(testinput)
	target := 0
	if result != target {
		t.Errorf("failed. got %d, should be %d\n", result, target)
	}
}
