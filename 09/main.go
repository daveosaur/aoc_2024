package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	// "strings"
	"time"
)

func part1(inp string) int {
	result := 0
	// var buff []string
	buff := make([]string, 0, 1000)
	counter := 0
	for i, c := range inp {
		num, _ := strconv.Atoi(string(c))
		switch i % 2 {
		case 0:
			num_to_write := strconv.Itoa(counter)
			for range num {
				buff = append(buff, num_to_write)
			}
			counter++
		case 1:
			for range num {
				buff = append(buff, ".")
			}
		}
	}
	//compact
	start := 0
	end := len(buff) - 1
	for {
		if start >= end {
			break
		}
		for buff[start] != "." && start < end {
			start++
		}
		for buff[end] == "." && end > start {
			end--
		}
		buff[start], buff[end] = buff[end], buff[start]
	}
	//count
	for i, s := range buff {
		num, _ := strconv.Atoi(s)
		result += num * i
	}

	return result
}

type file struct {
	ID         int
	Length     int
	WrittenLen int
	FreeSpace  int
	Seen       bool
}

func (f *file) String() string {
	var b strings.Builder
	num := strconv.Itoa(f.ID)
	for range f.Length {
		b.WriteString(num)
	}
	for range f.FreeSpace {
		b.WriteString(".")
	}
	return b.String()
}

func part2(inp string) int {
	// result := 0
	var result int = 0
	files := make([]*file, 0, 10000)
	id_counter := 0
	for i := 0; i < len(inp); i += 2 {
		// for i, c := range inp {
		file_len, _ := strconv.Atoi(string(inp[i]))
		//lazy whatever
		temp := strconv.Itoa(id_counter)

		free_space := func() int {
			if i+1 == len(inp) {
				return 0
			} else {
				n, _ := strconv.Atoi(string(inp[i+1]))
				return n
			}
		}()
		// free_space, _ := strconv.Atoi(string(inp[i+1]))
		files = append(files, &file{
			ID:         id_counter,
			Length:     file_len,
			WrittenLen: len(temp),
			FreeSpace:  free_space,
			Seen:       false,
		})
		id_counter++
	}
	//compact
	test := 0
	for i := len(files) - 1; i > 0; i-- {
		for j, f := range files {
			if j == i {
				break
			}
			if f.FreeSpace >= files[i].Length*files[i].WrittenLen && !files[i].Seen {
				files[i].Seen = true
				temp := files[i]
				files = slices.Delete(files, i, i+1)
				files = slices.Insert(files, j+1, temp)
				files[i].FreeSpace += temp.Length*temp.WrittenLen + temp.FreeSpace
				files[j+1].FreeSpace = f.FreeSpace - files[j+1].Length*files[j+1].WrittenLen
				f.FreeSpace = 0
				i++
				test++
				break
			}
		}
	}
	//count
	var buff strings.Builder
	for _, f := range files {
		buff.WriteString(f.String())
	}

	for i, s := range buff.String() {
		num, _ := strconv.Atoi(string(s))
		result += num * i
		if i%1000 == 0 {
			fmt.Println(result)
		}
	}
	// fmt.Println(buff.String())
	return result
}

func part2Fixed(inp string) int {
	// result := 0
	var result int = 0
	files := make([]*file, 0, 10000)
	id_counter := 0
	for i := 0; i < len(inp); i += 2 {
		// for i, c := range inp {
		file_len, _ := strconv.Atoi(string(inp[i]))
		//lazy whatever

		free_space := func() int {
			if i+1 == len(inp) {
				return 0
			} else {
				n, _ := strconv.Atoi(string(inp[i+1]))
				return n
			}
		}()
		// free_space, _ := strconv.Atoi(string(inp[i+1]))
		files = append(files, &file{
			ID:         id_counter,
			Length:     file_len,
			WrittenLen: 0,
			FreeSpace:  free_space,
			Seen:       false,
		})
		id_counter++
	}
	//compact
	test := 0
	for i := len(files) - 1; i > 0; i-- {
		for j, f := range files {
			if j == i {
				break
			}
			if f.FreeSpace >= files[i].Length && !files[i].Seen {
				files[i].Seen = true
				temp := files[i]
				files = slices.Delete(files, i, i+1)
				files = slices.Insert(files, j+1, temp)
				files[i].FreeSpace += temp.Length + temp.FreeSpace
				files[j+1].FreeSpace = f.FreeSpace - files[j+1].Length
				f.FreeSpace = 0
				i++
				test++
				break
			}
		}
	}
	//count
	var buff strings.Builder
	for _, f := range files {
		buff.WriteString(f.String())
	}
	index := 0
	for _, f := range files {
		for range f.Length {
			result += f.ID * index
			index++
		}
		index += f.FreeSpace
	}

	// for i, s := range buff.String() {
	// 	num, _ := strconv.Atoi(string(s))
	// 	result += num * i
	// 	if i%1000 == 0 {
	// 		fmt.Println(result)
	// 	}
	// }
	// fmt.Println(buff.String())
	return result
}

func main() {
	inp, _ := os.ReadFile("input.txt")
	start := time.Now()
	fmt.Printf("%d\n", part1(string(inp)))
	fmt.Printf("%d\ntotal time: %s\n", part2Fixed(string(inp)), time.Since(start).String())
}
