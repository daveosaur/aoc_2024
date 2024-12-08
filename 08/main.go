package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func part1(inp string) int {
	split := strings.Split(inp, "\n")
	split = split[:len(split)-1]
	locations := make(map[[2]int]bool)
	for i, line := range split {
		for j, c := range line {
			if c != '.' {
				//find the rest
				for sub_i, sub_line := range split {
					for sub_j, sub_c := range sub_line {
						if sub_i == i && sub_j == j {
							continue
						}
						if sub_c == c {
							//match found
							res1_i := i - (i-sub_i)*2
							res1_j := j - (j-sub_j)*2
							res2_i := sub_i + (i-sub_i)*2
							res2_j := sub_j + (j-sub_j)*2
							if res1_i >= 0 && res1_i < len(split) && res1_j >= 0 && res1_j < len(split[0]) {
								locations[[2]int{res1_i, res1_j}] = true
							}
							if res2_i >= 0 && res2_i < len(split) && res2_j >= 0 && res2_j < len(split[0]) {
								locations[[2]int{res2_i, res2_j}] = true
							}
						}
					}
				}
			}
		}
	}
	return len(locations)
}
func part2(inp string) int {
	split := strings.Split(inp, "\n")
	split = split[:len(split)-1]
	locations := make(map[[2]int]bool)
	for i, line := range split {
		for j, c := range line {
			if c != '.' {
				//find the rest
				for sub_i, sub_line := range split {
					for sub_j, sub_c := range sub_line {
						if sub_i == i && sub_j == j {
							continue
						}
						if sub_c == c {
							//match found
							j_dist := j - sub_j
							i_dist := i - sub_i
							temp_i, temp_j := i, j
							for {
								if temp_i < 0 || temp_i >= len(split) || temp_j < 0 || temp_j >= len(split[0]) {
									break
								}
								locations[[2]int{temp_i, temp_j}] = true
								temp_i -= i_dist
								temp_j -= j_dist
							}
							// i was literally running this twice for either direction, not counting on
							// running it from the other side on the next antenna match lol
						}
					}
				}
			}
		}
	}
	return len(locations)
}

func main() {
	inp, _ := os.ReadFile("input.txt")
	start := time.Now()
	fmt.Printf("%d\n", part1(string(inp)))
	fmt.Printf("%d\ntotal time: %s\n", part2(string(inp)), time.Since(start).String())
}
