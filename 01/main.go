package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
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
	inp, _ := os.ReadFile("input.txt")
	start := time.Now()
	fmt.Printf("%d\n", part1(string(inp)))
	fmt.Printf("%d\ntotal time: %s\n", part2(string(inp)), time.Since(start).String())
}
