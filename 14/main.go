package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type robot struct {
	x, y   int
	vx, vy int
}

// moves robot mod dimensions of the size of the room
func (r *robot) move(x_size, y_size int) {
	r.x = (r.x + r.vx) % x_size
	if r.x < 0 {
		r.x += x_size
	}
	r.y = (r.y + r.vy) % y_size
	if r.y < 0 {
		r.y += y_size
	}
}

func (r *robot) unmove(x_size, y_size int) {
	r.x = (r.x - r.vx) % x_size
	if r.x < 0 {
		r.x += x_size
	}
	r.y = (r.y - r.vy) % y_size
	if r.y < 0 {
		r.y += y_size
	}
}

func part1(inp string, x_size, y_size int) int {
	parser := regexp.MustCompile(`(-?\d+)`)
	positions := make(map[[2]int]int)
	for _, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			break
		}
		nums := parser.FindAllString(line, -1)
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		vx, _ := strconv.Atoi(nums[2])
		vy, _ := strconv.Atoi(nums[3])
		r := robot{x, y, vx, vy}
		for range 100 {
			r.move(x_size, y_size)
		}
		positions[[2]int{r.x, r.y}]++
	}
	for i := range y_size {
		for j := range x_size {
			if positions[[2]int{j, i}] > 0 {
				fmt.Printf("R")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	var nums [4]int
	middle := [2]int{(x_size / 2), (y_size / 2)}
	for pos, count := range positions {
		if pos[0] == middle[0] || pos[1] == middle[1] {
			continue
		}
		if pos[0] < middle[0] && pos[1] < middle[1] {
			nums[0] += count
		}
		if pos[0] < middle[0] && pos[1] > middle[1] {
			nums[1] += count
		}
		if pos[0] > middle[0] && pos[1] < middle[1] {
			nums[2] += count
		}
		if pos[0] > middle[0] && pos[1] > middle[1] {
			nums[3] += count
		}
	}
	return nums[0] * nums[1] * nums[2] * nums[3]
}

type state struct {
	robots         []*robot
	x_size, y_size int
	index          int
	count          int
	part1_result   int
}

// lol FINE lets make a viewer.
func (s *state) drawWindow() {
	if rl.IsKeyDown(rl.KeyRight) {
		s.count++
		if s.count > 3 {
			for _, r := range s.robots {
				r.move(s.x_size, s.y_size)
			}
			s.count = 0
			s.index++
		}
	} else if rl.IsKeyDown(rl.KeyLeft) {
		s.count++
		if s.count > 3 {
			for _, r := range s.robots {
				r.unmove(s.x_size, s.y_size)
			}
			s.count = 0
			s.index--
		}
	} else if rl.IsKeyDown(rl.KeyUp) {
		s.index += 10
		for _, r := range s.robots {
			for range 10 {
				r.move(s.x_size, s.y_size)
			}
		}
	} else if rl.IsKeyDown(rl.KeyDown) {
		s.index -= 10
		for _, r := range s.robots {
			for range 10 {
				r.unmove(s.x_size, s.y_size)
			}
		}
	}
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	for _, r := range s.robots {
		rl.DrawRectangle(int32(r.x*3), int32(r.y*3), 3, 3, rl.RayWhite)
	}
	rl.DrawText(fmt.Sprintf("iteration: %d", s.index), 300, 300, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("oh also part1 was: %d", s.part1_result), 300, 320, 20, rl.RayWhite)
	rl.EndDrawing()

}

func setup(inp string, x_size, y_size int) *state {
	s := state{make([]*robot, 0), x_size, y_size, 0, 0, 0}
	s.part1_result = part1(inp, x_size, y_size)
	parser := regexp.MustCompile(`(-?\d+)`)
	for _, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			break
		}
		nums := parser.FindAllString(line, -1)
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		vx, _ := strconv.Atoi(nums[2])
		vy, _ := strconv.Atoi(nums[3])
		s.robots = append(s.robots, &robot{x, y, vx, vy})
	}

	return &s
}

//yeah no. switched to raylib
/*
func part2(inp string, x_size, y_size int) int {
	// out, _ := os.Create("out.txt")
	// defer out.Close()
	parser := regexp.MustCompile(`(-?\d+)`)
	positions := make(map[[2]int]int)
	robots := make([]*robot, 0)
	for _, line := range strings.Split(inp, "\n") {
		if len(line) <= 0 {
			break
		}
		nums := parser.FindAllString(line, -1)
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		vx, _ := strconv.Atoi(nums[2])
		vy, _ := strconv.Atoi(nums[3])
		robots = append(robots, &robot{x, y, vx, vy})
		// for range 100 {
		// 	r.move(x_size, y_size)
		// }
		// positions[[2]int{r.x, r.y}]++
	}
	getline := bufio.NewReader(os.Stdin)
	// for n := range 100 {
	iter := 0
	for {
		iter++
		clear(positions)
		for _, r := range robots {
			r.move(x_size, y_size)
			positions[[2]int{r.x, r.y}]++
		}
		for i := range y_size {
			for j := range x_size {
				if positions[[2]int{j, i}] > 0 {
					// out.WriteString("R")
					fmt.Print("R")
				} else {
					fmt.Print(".")
					// out.WriteString(".")
				}
			}
			fmt.Println()
			// out.WriteString("\n")
		}
		fmt.Println(iter)
		getline.ReadString('\n')
		// out.WriteString(fmt.Sprintf("iteration %d\n", n))
	}
	// var nums [4]int
	// middle := [2]int{(x_size / 2), (y_size / 2)}
	// for pos, count := range positions {
	// 	if pos[0] == middle[0] || pos[1] == middle[1] {
	// 		continue
	// 	}
	// 	if pos[0] < middle[0] && pos[1] < middle[1] {
	// 		nums[0] += count
	// 	}
	// 	if pos[0] < middle[0] && pos[1] > middle[1] {
	// 		nums[1] += count
	// 	}
	// 	if pos[0] > middle[0] && pos[1] < middle[1] {
	// 		nums[2] += count
	// 	}
	// 	if pos[0] > middle[0] && pos[1] > middle[1] {
	// 		nums[3] += count
	// 	}
	// }
	// return nums[0] * nums[1] * nums[2] * nums[3]
	return 0
}
*/

func main() {
	inp, _ := os.ReadFile("input.txt")
	rl.InitWindow(640, 480, "bruh really")
	rl.SetTargetFPS(60)
	state := setup(string(inp), 101, 103)
	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {
		state.drawWindow()
	}
}
