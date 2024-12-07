package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParts(t *testing.T) {
	testinput := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	testinput2 := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	assert.Equal(t, part1(testinput), 161)
	assert.Equal(t, part2(testinput2), 48)

}
