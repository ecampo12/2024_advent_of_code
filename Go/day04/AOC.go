package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func (p *Point) add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}
func (p *Point) withinBounds(input []string) bool {
	return p.x >= 0 && p.x < len(input) && p.y >= 0 && p.y < len(input[0])
}

func part1(input []string) int {
	count := 0
	dirs := []Point{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	targets := "MAS"
	for r := range input {
		for c := range input[r] {
			if input[r][c] == 'X' {
				for _, dir := range dirs {
					p := Point{r, c}
					for k := 0; k < 3; k++ {
						p = p.add(dir)
						if !p.withinBounds(input) || input[p.x][p.y] != targets[k] {
							break
						}
						if k == 2 {
							count++
						}
					}
				}
			}
		}
	}

	return count
}

func part2(input []string) int {
	count := 0
	valid := map[string]struct{}{
		"MMSS": {}, "MSSM": {}, "SSMM": {}, "SMMS": {},
	}
	cornerOffsets := []Point{
		{-1, -1}, {-1, 1}, {1, 1}, {1, -1},
	}

	for r := 1; r < len(input)-1; r++ {
		for c := 1; c < len(input[r])-1; c++ {
			if input[r][c] == 'A' {
				corners := make([]byte, len(cornerOffsets))

				for i, offset := range cornerOffsets {
					corners[i] = input[r+offset.x][c+offset.y]
				}

				if _, ok := valid[string(corners)]; ok {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	part1_ans := part1(lines)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
