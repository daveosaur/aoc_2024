package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func part1(inp string) int {
	result := 0
	for _, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			break
		}
		parsed := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parsed[0])
		nums := make([]int, 0)
		for _, num := range strings.Split(parsed[1], " ") {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		if isValid(target, nums[0], nums[1:]) {
			result += target
		}
	}
	return result
}

func isValid(target int, acc int, inp []int) bool {
	if acc > target {
		return false
	}
	if len(inp) == 1 {
		return acc+inp[0] == target || acc*inp[0] == target
	}
	return isValid(target, acc+inp[0], inp[1:]) || isValid(target, acc*inp[0], inp[1:])
}

func isValidP2(target int, acc int, inp []int) bool {
	//small optimization
	if acc > target {
		return false
	}
	if len(inp) == 1 {
		return acc+inp[0] == target || acc*inp[0] == target || concatNums(acc, inp[0]) == target
	}
	return isValidP2(target, acc+inp[0], inp[1:]) || isValidP2(target, acc*inp[0], inp[1:]) || isValidP2(target, concatNums(acc, inp[0]), inp[1:])
}

// joining strings is dum
func concatNums(a, b int) int {
	temp := float32(b)
	for temp >= 1 {
		temp /= 10
		a *= 10
	}
	return a + b
}

func part2(inp string) int {
	result := 0
	for _, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			break
		}
		parsed := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parsed[0])
		nums := make([]int, 0, 10)
		for _, num := range strings.Split(parsed[1], " ") {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		if isValidP2(target, nums[0], nums[1:]) {
			result += target
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
