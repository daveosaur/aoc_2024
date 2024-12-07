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
	split := strings.Split(inp, "\n")
	parsingData := true
	rules := make(map[string][]string)
	for _, line := range split {
		if len(line) <= 0 {
			parsingData = false
			continue
		}
		if parsingData {
			nums := strings.Split(line, "|")
			rules[nums[0]] = append(rules[nums[0]], nums[1])
		} else {
			order := strings.Split(line, ",")
			for i := 0; i < len(order); i++ {
				for j := i; j < len(order); j++ {
					if slices.Contains(rules[order[j]], order[i]) {
						goto cancel
					}
				}
			}
			midNum, _ := strconv.Atoi(order[(len(order)-1)/2])
			result += midNum
		}
	cancel:
	}
	// fmt.Println(rules)
	return result
}

func part2(inp string) int {
	result := 0
	split := strings.Split(inp, "\n")
	parsingData := true
	rules := make(map[string][]string)
	swap := 0
	for _, line := range split {
		if len(line) <= 0 {
			parsingData = false
			continue
		}
		if parsingData {
			nums := strings.Split(line, "|")
			rules[nums[0]] = append(rules[nums[0]], nums[1])
		} else {
			illegalLine := false
			order := strings.Split(line, ",")
			for i := 0; i < len(order); i++ {
				for j := i; j < len(order); j++ {
					// if slices.Contains(rules[order[j]], order[i]) {
					// 	// fmt.Println(i, j, rules[order[j]])
					// 	fmt.Println(line)
					// 	goto loop
					// }
					if slices.Contains(rules[order[j]], order[i]) {
						order[i], order[j] = order[j], order[i]
						illegalLine = true
						swap++
					}
				}
			}
			if illegalLine {
				midNum, _ := strconv.Atoi(order[(len(order)-1)/2])
				result += midNum

			}
			continue
		}
	}
	return result
}

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println(part1(string(input)))
	fmt.Println(part2(string(input)))
}
