package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type pos [2]int

func part1(inp string) int {
	result := 0
	split := strings.Split(inp, "\n")
	split = split[:len(split)-1]
	for i := range split {
		for j := range split[0] {
			if split[i][j] == '0' {
				result += len(explorePos(split, j, i))
			}
		}
	}
	return result
}

func explorePos(inp []string, x, y int) map[pos]int {
	seen := make(map[pos]int)
	if y > 0 && inp[y-1][x] == '1' {
		subExplore(inp, x, y-1, seen)
	}
	if y < len(inp)-1 && inp[y+1][x] == '1' {
		subExplore(inp, x, y+1, seen)
	}
	if x > 0 && inp[y][x-1] == '1' {
		subExplore(inp, x-1, y, seen)
	}
	if x < len(inp[0])-1 && inp[y][x+1] == '1' {
		subExplore(inp, x+1, y, seen)
	}
	return seen
}

func subExplore(inp []string, x, y int, seen map[pos]int) {
	c := inp[y][x]
	if c == '9' {
		seen[[2]int{y, x}]++
		return
	}
	if y > 0 && inp[y-1][x] == c+1 {
		subExplore(inp, x, y-1, seen)
	}
	if y < len(inp)-1 && inp[y+1][x] == c+1 {
		subExplore(inp, x, y+1, seen)
	}
	if x > 0 && inp[y][x-1] == c+1 {
		subExplore(inp, x-1, y, seen)
	}
	if x < len(inp[0])-1 && inp[y][x+1] == c+1 {
		subExplore(inp, x+1, y, seen)
	}
}

func part2(inp string) int {
	result := 0
	split := strings.Split(inp, "\n")
	split = split[:len(split)-1]
	for i := range split {
		for j := range split[0] {
			if split[i][j] == '0' {
				seen := explorePos(split, j, i)
				for _, paths := range seen {
					result += paths
				}
			}
		}
	}
	return result
}

func main() {
	inp, _ := os.ReadFile("input.txt")
	start_total := time.Now()
	fmt.Printf("part 1: %d - %s\n", part1(string(inp)), time.Since(start_total))
	start_p2 := time.Now()
	fmt.Printf("part 2: %d - %s\n", part2(string(inp)), time.Since(start_p2))
	fmt.Printf("total time: %s\n", time.Since(start_total))
}
