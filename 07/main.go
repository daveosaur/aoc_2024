package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(inp string) int {
	result := 0
	for _, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			break
		}
		parsed := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parsed[0])
		nums_s := strings.Split(parsed[1], " ")
		//gather up the numbers
		nums := make([]int, 0)
		for _, num := range nums_s {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		if isValid(target, 0, nums) {
			result += target
		}
		// fmt.Println(result, ": ", line)
	}
	return result
}

func isValid(target int, acc int, inp []int) bool {
	if len(inp) == 1 {
		if acc+inp[0] == target || acc*inp[0] == target {
			return true
		}
	}
	if len(inp) > 1 {
		if isValid(target, acc+inp[0], inp[1:]) {
			return true
		}
		if acc == 0 {
			return isValid(target, inp[0], inp[1:])
		} else {
			return isValid(target, acc*inp[0], inp[1:])

		}
	}
	return false
}

func isValidP2(target int, acc int, inp []int) bool {
	if len(inp) == 0 {
		return acc == target
	}
	if len(inp) == 1 {
		if acc+inp[0] == target || acc*inp[0] == target {
			return true
		}
		acc_s := strconv.Itoa(acc)
		inp_first_s := strconv.Itoa(inp[0])
		merged, _ := strconv.Atoi(acc_s + inp_first_s)
		return merged == target

	}
	if len(inp) > 1 {
		if isValidP2(target, acc+inp[0], inp[1:]) {
			return true
		}
		if acc == 0 {
			if isValidP2(target, inp[0], inp[1:]) {
				return true
			}
			inp_first_s := strconv.Itoa(inp[0])
			inp_sec_s := strconv.Itoa(inp[1])
			merged, _ := strconv.Atoi(inp_first_s + inp_sec_s)
			if isValidP2(target, merged, inp[2:]) {
				return true
			}
		} else {
			if isValidP2(target, acc*inp[0], inp[1:]) {
				return true
			}
			acc_s := strconv.Itoa(acc)
			inp_first_s := strconv.Itoa(inp[0])
			merged, _ := strconv.Atoi(acc_s + inp_first_s)
			if isValidP2(target, merged, inp[1:]) {
				return true
			}
		}
	}
	return false
}

func part2(inp string) int {
	result := 0
	for _, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			break
		}
		parsed := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parsed[0])
		nums_s := strings.Split(parsed[1], " ")
		//gather up the numbers
		nums := make([]int, 0)
		for _, num := range nums_s {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		if isValidP2(target, 0, nums) {
			result += target
			// fmt.Println(result, ": ", line, " success!")
		} else {
			// fmt.Println(result, ": ", line)

		}
	}
	return result
}

func main() {
	inp, _ := os.ReadFile("input.txt")
	fmt.Println(part1(string(inp)))
	fmt.Println(part2(string(inp)))
}
