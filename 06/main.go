package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type direction int

const (
	NONE direction = iota
	UP
	RIGHT
	DOWN
	LEFT
)

type bump struct {
	x, y int
}

func part1(inp string) int {
	count := 0
	split := strings.Split(inp, "\n")
	split = split[:len(split)-1]
	posx, posy := 0, 0
	dir := UP
	seen := make(map[[2]int]bool)
	for i := range split {
		for j := range split[i] {
			if split[i][j] == '^' {
				posx = j
				posy = i
				goto cont
			}
		}
	}
cont:
	for {
		count++
		seen[[2]int{posy, posx}] = true
		switch dir {
		case UP:
			if posy == 0 {
				goto done
			}
			if split[posy-1][posx] == '#' {
				dir = RIGHT
			} else {
				posy--
			}
		case RIGHT:
			if posx == len(split[0])-1 {
				goto done
			}
			if split[posy][posx+1] == '#' {
				dir = DOWN
			} else {
				posx++
			}
		case DOWN:
			if posy == len(split)-1 {
				goto done
			}
			if split[posy+1][posx] == '#' {
				dir = LEFT
			} else {
				posy++
			}
		case LEFT:
			if posx == 0 {
				goto done
			}
			if split[posy][posx-1] == '#' {
				dir = UP
			} else {
				posx--
			}
		default: //something fucked up
		}
	}
done:
	return len(seen)
}

// not actually working. bleh
func part2(inp string) int {
	result := 0

	split := strings.Split(inp, "\n")
	split = split[:len(split)-1]
	posx, posy := 0, 0
	dir := UP
	loopPoints := make(map[[2]int][]direction)
	for i := range split {
		for j := range split[i] {
			if split[i][j] == '^' {
				posx = j
				posy = i
				goto cont
			}
		}
	}
cont:
	//now build up the map for all the loop points
	for i, line := range split {
		for j, c := range line {
			if c == '#' {
				//fill em up!!!
				//up(actually down)
				for pos_i := i + 1; pos_i < len(split); pos_i++ {
					if split[pos_i][j] == '#' {
						break
					}
					loopPoints[[2]int{pos_i, j}] = append(loopPoints[[2]int{pos_i, j}], LEFT)
				}
				//down
				for pos_i := i - 1; pos_i > 0; pos_i-- {
					if split[pos_i][j] == '#' {
						break
					}
					loopPoints[[2]int{pos_i, j}] = append(loopPoints[[2]int{pos_i, j}], RIGHT)
				}
				//right(actually left)
				for pos_j := j - 1; pos_j > 0; pos_j-- {
					if split[i][pos_j] == '#' {
						break
					}
					loopPoints[[2]int{i, pos_j}] = append(loopPoints[[2]int{i, pos_j}], UP)
				}
				//left
				for pos_j := j + 1; pos_j < len(split[0])-1; pos_j++ {
					if split[i][pos_j] == '#' {
						break
					}
					loopPoints[[2]int{i, pos_j}] = append(loopPoints[[2]int{i, pos_j}], DOWN)
				}
			}

		}
	}
	// fmt.Println(loopPoints)

	//rerun the loop with loop point data i guess
	for {
		switch dir {
		case UP:
			if posy == 0 {
				goto actuallydone
			}
			if split[posy-1][posx] == '#' {
				dir = RIGHT
			} else {
				posy--
			}
		case RIGHT:
			if posx == len(split[0])-1 {
				goto actuallydone
			}
			if split[posy][posx+1] == '#' {
				dir = DOWN
			} else {
				posx++
			}
		case DOWN:
			if posy == len(split)-1 {
				goto actuallydone
			}
			if split[posy+1][posx] == '#' {
				dir = LEFT
			} else {
				posy++
			}
		case LEFT:
			if posx == 0 {
				goto actuallydone
			}
			if split[posy][posx-1] == '#' {
				dir = UP
			} else {
				posx--
			}
		default: //something fucked up
		}

		loopDir, ok := loopPoints[[2]int{posy, posx}]
		if ok {
			if slices.Contains(loopDir, dir) && checkForLoop(posx, posy, dir, split) {
				// fmt.Println(posy, posx)
				result++
			}
		}
	}
actuallydone:
	return result
}

func checkForLoop(posx, posy int, d direction, split []string) bool {
	startx := posx
	starty := posy
	startd := d
	switch d {
	case UP:
		d = RIGHT
	case RIGHT:
		d = DOWN
	case DOWN:
		d = LEFT
	case LEFT:
		d = UP
	}

	count := 0
	for {
		count++
		//hard cap timeout. if you hit this you're probably looping!
		if count > 1000 {
			return true
		}
		switch d {
		case UP:
			if posy == 0 {
				return false
			}
			if split[posy-1][posx] == '#' {
				d = RIGHT
			} else {
				posy--
			}
		case RIGHT:
			if posx == len(split[0])-1 {
				return false
			}
			if split[posy][posx+1] == '#' {
				d = DOWN
			} else {
				posx++
			}
		case DOWN:
			if posy == len(split)-1 {
				return false
			}
			if split[posy+1][posx] == '#' {
				d = LEFT
			} else {
				posy++
			}
		case LEFT:
			if posx == 0 {
				return false
			}
			if split[posy][posx-1] == '#' {
				d = UP
			} else {
				posx--
			}
		default: //something fucked up
		}
		if posx == startx && posy == starty && d == startd {
			return true
		}
	}
}

func checkForLoopArr(posx, posy int, d direction, split [][]rune) bool {
	startx := posx
	starty := posy
	startd := d

	count := 0
	for {
		count++
		//hard cap timeout. if you hit this you're probably looping!
		if count > 10000 {
			return true
		}
		switch d {
		case UP:
			if posy == 0 {
				return false
			}
			if split[posy-1][posx] == '#' {
				d = RIGHT
			} else {
				posy--
			}
		case RIGHT:
			if posx == len(split[0])-1 {
				return false
			}
			if split[posy][posx+1] == '#' {
				d = DOWN
			} else {
				posx++
			}
		case DOWN:
			if posy == len(split)-1 {
				return false
			}
			if split[posy+1][posx] == '#' {
				d = LEFT
			} else {
				posy++
			}
		case LEFT:
			if posx == 0 {
				return false
			}
			if split[posy][posx-1] == '#' {
				d = UP
			} else {
				posx--
			}
		default: //something fucked up
		}
		if posx == startx && posy == starty && d == startd {
			return true
		}
	}
}

// this is so stupid i hate myself
func part2BruteForce(inp string) int {
	result := 0
	split := make([][]rune, 1)
	startx, starty := 0, 0
	tempwhatevershutup := 0
	for _, c := range inp {
		switch c {
		case '\n':
			split = append(split, make([]rune, 0))
			tempwhatevershutup++
		default:
			split[tempwhatevershutup] = append(split[tempwhatevershutup], c)
		}
	}
	split = split[:len(split)-1]

	for i := range split {
		for j := range split[i] {
			if split[i][j] == '^' {
				startx = j
				starty = i
				goto cont
			}
		}
	}
cont:
	for i := range split {
		for j := range split[0] {
			if split[i][j] != '.' {
				continue
			}
			split[i][j] = '#'
			if checkForLoopArr(startx, starty, UP, split) {
				result++
			}
			split[i][j] = '.'
		}
	}
	return result

}

func main() {
	inp, _ := os.ReadFile("input.txt")
	fmt.Println(part1(string(inp)))
	// fmt.Println(part2(string(inp)))
	fmt.Println(part2BruteForce(string(inp)))
}
