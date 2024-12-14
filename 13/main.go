package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type pair struct {
	x, y int
}

func part1(inp string) int {
	result := 0
	data := make([]pair, 0, 3)
	for i, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			for i := range 100 {
				for j := range 100 {
					if i*data[0].x+j*data[1].x == data[2].x && i*data[0].y+j*data[1].y == data[2].y {
						result += 3*i + j
					}
				}
			}
			data = make([]pair, 0, 3)
			continue
		}
		_ = i
		dat := strings.Split(line, ": ")
		nums := strings.Split(dat[1], ", ")
		btn_x, _ := strconv.Atoi(nums[0][2:])
		btn_y, _ := strconv.Atoi(nums[1][2:])
		//a will be [0], b will be [1], prize is [2]
		data = append(data, pair{btn_x, btn_y})

	}
	return result
}

func part2(inp string) int {

	result := 0
	data := make([]pair, 0, 3)
	for i, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			data[2].x += 10000000000000
			data[2].y += 10000000000000
			for i := 10000000000000; i < 10000000000000+100; i++ {
				for j := 10000000000000; j < 10000000000000+100; j++ {
					fmt.Println(i, j)
					// for i := range 100 {
					// for j := range 100 {
					if i*data[0].x+j*data[1].x == data[2].x && i*data[0].y+j*data[1].y == data[2].y {
						result += 3*i + j
					}
				}
			}
			data = make([]pair, 0, 3)
			continue
		}
		_ = i
		dat := strings.Split(line, ": ")
		nums := strings.Split(dat[1], ", ")
		btn_x, _ := strconv.Atoi(nums[0][2:])
		btn_y, _ := strconv.Atoi(nums[1][2:])
		//a will be [0], b will be [1], prize is [2]
		data = append(data, pair{btn_x, btn_y})

	}
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
