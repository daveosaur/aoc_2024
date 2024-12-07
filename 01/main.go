package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(inp string) int {
	result := 0
	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)
	for _, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			break
		}
		nums := strings.Split(line, "   ")
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		left = append(left, l)
		right = append(right, r)
	}
	slices.Sort(left)
	slices.Sort(right)
	for i := range left {
		result += max(left[i], right[i]) - min(left[i], right[i])
	}
	return result
}

func part2(inp string) int {
	result := 0
	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)
	for _, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			break
		}
		nums := strings.Split(line, "   ")
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		left = append(left, l)
		right = append(right, r)
	}
	//small optimization i guess
	slices.Sort(right)
	for _, lnum := range left {
		for _, rnum := range right {
			if lnum < rnum {
				break
			}
			if lnum == rnum {
				result += lnum
			}
		}
	}
	return result
}

func main() {
	testinput := `3   4
4   3
2   5
1   3
3   9
3   3
`
	_ = testinput
	// fmt.Println(part1(testinput))
	// fmt.Println(part2(testinput))
	inp, _ := os.ReadFile("input.txt")
	_ = inp
	fmt.Println(part1(string(inp)))
	fmt.Println(part2(string(inp)))
}
