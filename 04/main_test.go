package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	testinput := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	assert.Equal(t, 18, part1(testinput), "they should be equal??")

}

func TestPart2(t *testing.T) {
	testinput := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........
`

	assert.Equal(t, 9, part2(testinput))
}
