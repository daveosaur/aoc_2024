package main

import (
	"fmt"
	"os"
	"strings"
)

func part1(inp string) int {
	result := 0
	split := strings.Split(inp, "\n")
	split = split[:len(split)-1]
	for i, line := range split {
		for j := range line {
			if split[i][j] == 'X' {
				result += checkForXmas(j, i, split)
			}

		}
	}
	return result
}

// checks for the number of XMAS found from position
func checkForXmas(x, y int, inp []string) int {
	result := 0
	//check horizontal
	if x+3 < len(inp[y]) && inp[y][x:x+4] == "XMAS" {
		result++
	}
	if x-3 >= 0 && inp[y][x-3:x+1] == "SAMX" {
		result++
	}
	//vertical
	if y+3 < len(inp) {
		if inp[y+1][x] == 'M' && inp[y+2][x] == 'A' && inp[y+3][x] == 'S' {
			result++
		}
	}
	if y-3 >= 0 {
		if inp[y-1][x] == 'M' && inp[y-2][x] == 'A' && inp[y-3][x] == 'S' {
			result++
		}
	}
	//diagonals
	if y+3 < len(inp) && x+3 < len(inp[y]) {
		if inp[y+1][x+1] == 'M' && inp[y+2][x+2] == 'A' && inp[y+3][x+3] == 'S' {
			result++
		}
	}
	if y+3 < len(inp) && x-3 >= 0 {
		if inp[y+1][x-1] == 'M' && inp[y+2][x-2] == 'A' && inp[y+3][x-3] == 'S' {
			result++
		}
	}
	if y-3 >= 0 && x+3 < len(inp[y]) {
		if inp[y-1][x+1] == 'M' && inp[y-2][x+2] == 'A' && inp[y-3][x+3] == 'S' {
			result++
		}
	}
	if y-3 >= 0 && x-3 >= 0 {
		if inp[y-1][x-1] == 'M' && inp[y-2][x-2] == 'A' && inp[y-3][x-3] == 'S' {
			result++
		}
	}

	return result
}

func checkForMasX(x, y int, inp []string) bool {
	if x+3 > len(inp[y]) || y+3 > len(inp) || inp[y+1][x+1] != 'A' {
		return false
	}
	//check diagonal downforward
	if inp[y][x] == 'M' && inp[y+2][x+2] == 'S' {
		if inp[y][x+2] == 'M' && inp[y+2][x] == 'S' {
			return true
		}
		if inp[y+2][x] == 'M' && inp[y][x+2] == 'S' {
			return true
		}
		return false
	}
	//diagonal downback
	if inp[y][x+2] == 'M' && inp[y+2][x] == 'S' {
		if inp[y][x] == 'S' && inp[y+2][x+2] == 'M' {
			return true
		}
		return false
	}
	//diagonal upforward
	if inp[y+2][x] == 'M' && inp[y][x+2] == 'S' {
		if inp[y+2][x+2] == 'M' && inp[y][x] == 'S' {
			return true
		}
	}
	return false
}

func part2(inp string) int {
	result := 0
	split := strings.Split(inp, "\n")
	split = split[:len(split)-1]
	for i, line := range split {
		for j := range line {
			if checkForMasX(j, i, split) {
				result++
			}
		}
	}
	return result
}

func main() {
	inp, _ := os.ReadFile("input.txt")
	fmt.Println("part 1: ", part1(string(inp)))
	fmt.Println("part 2: ", part2(string(inp)))
}
