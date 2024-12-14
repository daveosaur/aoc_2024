package main

import (
	"testing"
)

var testinput string = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`

func TestP1(t *testing.T) {

	result := part1(testinput, 11, 7)
	target := 12
	if result != target {
		t.Errorf("failed. got %d, should be %d\n", result, target)
	}
}

// func TestP2(t *testing.T) {

// 	result := part2(testinput)
// 	target := 0
// 	if result != target {
// 		t.Errorf("failed. got %d, should be %d\n", result, target)
// 	}
// }
