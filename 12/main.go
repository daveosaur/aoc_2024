package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type dir int

const (
	UP dir = iota
	RIGHT
	DOWN
	LEFT
)

func part1(inp string) int {
	result := 0
	split := strings.Split(inp, "\n")
	split = split[:len(split)-1]
	seen := make(map[[2]int]bool)
	for i := range split {
		for j := range split[0] {
			temp := len(seen)
			_, ok := seen[[2]int{i, j}]
			if !ok {
				perimiter := calcRegion(split[i][j], j, i, split, seen)
				area := len(seen) - temp
				fence := perimiter * area
				// fmt.Println(perimiter, fence, area)
				result += fence
			}
		}
	}

	return result
}

func calcRegion(plant byte, x, y int, inp []string, m map[[2]int]bool) int {
	if x < 0 || x >= len(inp[0]) || y < 0 || y >= len(inp) {
		return 1 //oob, which means this is an edge and has a perimiter wall
	}
	if inp[y][x] != plant {
		return 1 //edge, 1 perimiter
	}
	if m[[2]int{y, x}] == true {
		return 0 //not edge, already seen.
	}
	//else
	m[[2]int{y, x}] = true
	return calcRegion(plant, x-1, y, inp, m) + calcRegion(plant, x+1, y, inp, m) + calcRegion(plant, x, y-1, inp, m) + calcRegion(plant, x, y+1, inp, m)
}

func part2(inp string) int {
	result := 0
	// split := strings.Split(inp, "\n")
	// split = split[:len(split)-1]
	// seen := make(map[[2]int]byte)
	// for i := range split {
	// 	for j := range split[0] {
	// 		temp := len(seen)
	// 		_, ok := seen[[2]int{i, j}]
	// 		if !ok {
	// 			perimiter := calcRegionP2(UP, split[i][j], j, i, split, seen)
	// 			area := len(seen) - temp
	// 			fence := perimiter * area
	// 			fmt.Println(area, perimiter)
	// 			result += fence
	// 		}
	// 	}
	// }

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
