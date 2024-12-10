package main

import (
	"testing"
)

var testinput string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

func TestP1(t *testing.T) {

	result := part1(testinput)
	target := 36
	if result != target {
		t.Errorf("failed. got %d, should be %d\n", result, target)
	}
}
func TestP2(t *testing.T) {

	result := part2(testinput)
	target := 81
	if result != target {
		t.Errorf("failed. got %d, should be %d\n", result, target)
	}
}
