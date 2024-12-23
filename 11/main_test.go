package main

import (
	"testing"
)

var testinput string = `125 17`

func TestP1(t *testing.T) {

	result := part1(testinput)
	target := 55312
	if result != target {
		t.Errorf("failed. got %d, should be %d\n", result, target)
	}
}

func TestP2(t *testing.T) {

	result := solveMemo(25, testinput)
	target := 55312
	if result != target {
		t.Errorf("failed. got %d, should be %d\n", result, target)
	}
}
