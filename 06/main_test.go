package main

import (
	"testing"
)

func TestP1(t *testing.T) {
	testinput := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

	result := part1(testinput)
	if result != 41 {
		t.Errorf("failed. got %d, should be 41", result)
	}
}

func TestP2(t *testing.T) {
	testinput := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

	result := part2(testinput)
	if result != 6 {
		t.Errorf("failed. got %d, should be 6", result)
	}
}

func TestBrute(t *testing.T) {
	testinput := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

	result := part2BruteForce(testinput)
	if result != 6 {
		t.Errorf("failed. got %d, should be 6", result)
	}
}