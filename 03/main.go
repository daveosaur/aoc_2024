package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1(inp string) int {
	result := 0
	reg := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	split := reg.FindAllStringSubmatch(inp, -1)

	for _, s := range split {
		// fmt.Println(s)
		first, _ := strconv.Atoi(s[1])
		second, _ := strconv.Atoi(s[2])
		result += first * second
	}
	return result
}

func part2(inp string) int {
	//two regex because i suck and i'm tired.
	findDont := regexp.MustCompile(`don't\(\)(?:.|\s)*?do\(\)`)
	finish := regexp.MustCompile(`don't\(\).*`)

	modified := findDont.ReplaceAllString(inp, "")
	finished := finish.ReplaceAllString(modified, "")
	return part1(finished)

}

func main() {
	testinput := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	_ = testinput
	testinput2 := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	_ = testinput2

	// fmt.Println(part1(testinput))
	inp, _ := os.ReadFile("input.txt")
	_ = inp
	fmt.Println(part1(string(inp)))
	// fmt.Println(part2(testinput2))
	fmt.Println(part2(string(inp)))

}
