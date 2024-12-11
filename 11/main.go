package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

// stuff for when i was trying to actually brute force lol
var mut sync.Mutex
var p2count int

func part1(inp string) int {
	stones := make([]int, 0, 100)
	for _, s := range strings.Split(inp, " ") {
		num, _ := strconv.Atoi(s)
		stones = append(stones, num)
	}
	for range 25 {
		stones = parseStones(stones)
	}
	return len(stones)
}

func parseStones(inp []int) []int {
	for i := 0; ; i++ {
		if i >= len(inp) {
			break
		}
		if inp[i] == 0 {
			inp[i]++
			continue
		}
		left, right, s_len := getDigitsLength(inp[i])
		if s_len%2 == 0 {
			inp = slices.Delete(inp, i, i+1)
			inp = slices.Insert(inp, i, left)
			inp = slices.Insert(inp, i+1, right)
			i += 1
			continue
		}
		inp[i] *= 2024

	}
	return inp
}

func parseStonesP2(depth int, inp int) {
	if depth == 75 {
		mut.Lock()
		p2count++
		mut.Unlock()
		return
	}
	if inp == 0 {
		inp++
		parseStonesP2(depth+1, inp)
		return
	} else {
		left, right, length := getDigitsLength(inp)
		// stone_str := strconv.Itoa(inp.val)
		// s_len := len(stone_str)
		if length%2 == 0 {
			parseStonesP2(depth+1, left)
			parseStonesP2(depth+1, right)
			return
		}
		//else, mult
		parseStonesP2(depth+1, inp*2024)
		return
	}
}

func getDigitsLength(inp int) (int, int, int) {
	temp := inp
	length := 0
	for temp > 0 {
		temp /= 10
		length++
	}
	left := inp
	right := 0
	multby := 1
	for range length / 2 {
		right += (left % 10) * multby
		multby *= 10
		left /= 10
	}
	return left, right, length
}

// bad bad bad no good
func part2(inp string) int {
	stones := make([]int, 0, 100)
	for _, s := range strings.Split(inp, " ") {
		num, _ := strconv.Atoi(s)
		stones = append(stones, num)
	}
	var wg sync.WaitGroup
	for _, s := range stones {
		wg.Add(1)
		go func() {
			parseStonesP2(0, s)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(p2count)
	return p2count
}

// global map for memoization
var parseMap = map[[2]int]int{}

func recurseParseStones(inp int, depth int) int {
	if depth == 75 {
		temp := parseSingleStone(inp)
		return len(temp)
	}
	result, ok := parseMap[[2]int{inp, depth}]
	if !ok {
		stones := parseSingleStone(inp)
		for _, s := range stones {
			parseMap[[2]int{s, depth + 1}] = recurseParseStones(s, depth+1)
			result += parseMap[[2]int{s, depth + 1}]
		}
	}
	return result
}

func parseSingleStone(inp int) []int {
	result := []int{}
	if inp == 0 {
		result = append(result, 1)
	} else {
		left, right, s_len := getDigitsLength(inp)
		if s_len%2 == 0 {
			result = append(result, left, right)
		} else {
			result = append(result, inp*2024)
		}
	}
	return result
}

// the actual solution
func part2Memo(inp string) int {
	result := 0
	stones := make([]int, 0)
	for _, s := range strings.Split(inp, " ") {
		num, _ := strconv.Atoi(s)
		stones = append(stones, num)
	}
	for _, s := range stones {
		result += recurseParseStones(s, 1)
	}
	return result
}

func main() {
	inp, _ := os.ReadFile("input.txt")
	start_total := time.Now()
	fmt.Printf("part 1: %d - %s\n", part1(string(inp)), time.Since(start_total))
	start_p2 := time.Now()
	fmt.Printf("part 2: %d - %s\n", part2Memo(string(inp)), time.Since(start_p2).String())
	fmt.Printf("total time: %s\n", time.Since(start_total))
}
