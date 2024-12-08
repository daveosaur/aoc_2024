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
	start := time.Now()
	fmt.Printf("%d\n", part1(string(inp)))
	fmt.Printf("%d\ntotal time: %s\n", part2(string(inp)), time.Since(start).String())
}
