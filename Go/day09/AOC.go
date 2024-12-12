package main

import (
	"fmt"
	"os"
	"slices"
)

func removeTrailingFreeSpace(disk []int) []int {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == -1 {
			disk = disk[:i]
		} else {
			break
		}
	}
	return disk
}

func part1(input string) int {
	disk := []int{}
	fid := 0
	for c, char := range input {
		size := int(char - '0')
		if c%2 == 0 {
			disk = append(disk, slices.Repeat([]int{fid}, size)...)
			fid++
		} else {
			disk = append(disk, slices.Repeat([]int{-1}, size)...)
		}
	}

	freeSpace := []int{}
	for i, block := range disk {
		if block == -1 {
			freeSpace = append(freeSpace, i)
		}
	}

	for _, space := range freeSpace {
		if disk[len(disk)-1] == -1 {
			disk = removeTrailingFreeSpace(disk)
		}
		if len(disk) <= space {
			break
		}
		disk[space] = disk[len(disk)-1]
		disk = disk[:len(disk)-1]
	}
	sum := 0
	for i, block := range disk {
		sum += i * block
	}
	return sum
}

type Block struct {
	pos, size int
}

func part2(input string) int {
	files := map[int]Block{}
	freeSpace := []Block{}
	fid, pos := 0, 0

	for c, char := range input {
		size := int(char - '0')
		if c%2 == 0 {
			files[fid] = Block{pos, size}
			fid++
		} else {
			freeSpace = append(freeSpace, Block{pos, size})
		}
		pos += size
	}

	for id := fid - 1; id >= 0; id-- {
		file := files[id]
		for s, space := range freeSpace {
			if space.pos >= file.pos {
				freeSpace = freeSpace[:s]
				break
			}
			if file.size <= space.size {
				files[id] = Block{space.pos, file.size}
				if file.size == space.size {
					freeSpace = append(freeSpace[:s], freeSpace[s+1:]...)
				} else {
					freeSpace[s] = Block{space.pos + file.size, space.size - file.size}
				}
				break
			}
		}
	}

	sum := 0
	for i, block := range files {
		for j := block.pos; j < block.pos+block.size; j++ {
			sum += i * j
		}
	}
	return sum
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := string(input)
	part1_ans := part1(lines)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
