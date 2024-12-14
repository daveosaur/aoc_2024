package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

// global map for memoization
var parseMap = map[[2]int]int{}

// kept because its funny
func part1(inp string) int {
	stones := make([]int, 0, 100)
	for _, s := range strings.Split(inp, " ") {
		num, _ := strconv.Atoi(s)
		stones = append(stones, num)
	}
	for range 25 {
		stones = parseStonesP1(stones)
	}
	return len(stones)
}

func parseStonesP1(inp []int) []int {
	for i := 0; ; i++ {
		if i >= len(inp) {
			break
		}
		if inp[i] == 0 {
			inp[i]++
			continue
		}
		left, right, s_len := getDigitsLength(inp[i])
		if s_len%2 == 0 {
			inp = slices.Delete(inp, i, i+1)
			inp = slices.Insert(inp, i, left)
			inp = slices.Insert(inp, i+1, right)
			i += 1
			continue
		}
		inp[i] *= 2024

	}
	return inp
}

func getDigitsLength(inp int) (int, int, int) {
	temp := inp
	length := 0
	for temp > 0 {
		temp /= 10
		length++
	}
	left := inp
	right := 0
	multby := 1
	for range length / 2 {
		right += (left % 10) * multby
		multby *= 10
		left /= 10
	}
	return left, right, length
}

// now returns left, right, count instead of a slice
func parseSingleStone(inp int) (int, int, int) {
	if inp == 0 {
		return 1, 0, 1
	} else {
		left, right, s_len := getDigitsLength(inp)
		if s_len%2 == 0 {
			return left, right, 2
		} else {
			return inp * 2024, 0, 1
		}
	}
	// return result
}

func parseStones(target int, inp int, depth int) int {
	if depth == target {
		_, _, length := parseSingleStone(inp)
		return length
	}
	result, ok := parseMap[[2]int{inp, depth}]
	if !ok {
		left, right, length := parseSingleStone(inp)
		switch length {
		case 2:
			parseMap[[2]int{right, depth + 1}] = parseStones(target, right, depth+1)
			result += parseMap[[2]int{right, depth + 1}]
			fallthrough
		case 1:
			parseMap[[2]int{left, depth + 1}] = parseStones(target, left, depth+1)
			result += parseMap[[2]int{left, depth + 1}]
		default:
			//something fucked up
			panic("wat")
		}
	}
	return result
}

// the actual solution
func solveMemo(target int, inp string) int {
	clear(parseMap)
	result := 0
	for _, s := range strings.Split(inp, " ") {
		num, _ := strconv.Atoi(s)
		result += parseStones(target, num, 1)
	}
	return result
}

func main() {
	inp, _ := os.ReadFile("input.txt")
	start_total := time.Now()
	fmt.Printf("part 1: %d - %s\n", solveMemo(25, string(inp)), time.Since(start_total))
	start_p2 := time.Now()
	fmt.Printf("part 2: %d - %s\n", solveMemo(75, string(inp)), time.Since(start_p2).String())
	fmt.Printf("total time: %s\n", time.Since(start_total))
}
