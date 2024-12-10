package main

import (
	"fmt"
	"os"
	"time"
)

func part1(inp string) int {
	result := 0
	return result
}
func part2(inp string) int {
	result := 0
	return result
}

func main() {
	inp, _ := os.ReadFile("input.txt")
	start_total := time.Now()
	fmt.Printf("part 1: %d - %s\n", part1(string(inp)), time.Since(start_total))
	start_p2 := time.Now()
	fmt.Printf("part 2: %d - %s\n", part2(string(inp)), time.Since(start_p2).String())
	fmt.Printf("total time: %s\n", time.Since(start_total))
}
