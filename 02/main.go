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
	for _, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			break
		}
		nums := strings.Split(line, " ")
		// fmt.Println(nums)
		if isSafe(nums) {
			result++
		}
	}

	return result
}

func isSafe(nums []string) bool {
	last := 0
	var increasing bool
	for i, num := range nums {
		n, _ := strconv.Atoi(num)
		if i == 0 {
			last = n
		} else if i == 1 {
			if n > last && n-last <= 3 {
				increasing = true
			} else if n < last && last-n <= 3 {
				increasing = false
			} else {
				//unsafe
				return false
			}
			last = n
		} else {
			if n > last && increasing == true && (n-last <= 3) {
				last = n
			} else if n < last && increasing == false && (last-n <= 3) {
				last = n
			} else {
				//unsafe
				return false
			}
		}
	}
	return true
}

func part2(inp string) int {
	result := 0
	for _, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			break
		}
		nums := strings.Split(line, " ")
		if isSafe(nums) {
			result++
		} else {
			for i := range nums {
				tempNums := slices.Delete(slices.Clone(nums), i, i+1)
				if isSafe(tempNums) {
					result++
					break
				}
			}
		}
	}

	return result
}

func main() {
	inp, _ := os.ReadFile("input.txt")
	fmt.Println(part1(string(inp)))
	fmt.Println(part2(string(inp)))

}
