package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
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
	inp, _ := os.ReadFile("input.txt")
	start := time.Now()
	fmt.Printf("%d\n", part1(string(inp)))
	fmt.Printf("%d\ntotal time: %s\n", part2(string(inp)), time.Since(start).String())
}
