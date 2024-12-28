package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct{ row, col int }

func add(a, b Point) Point { return Point{a.row + b.row, a.col + b.col} }

func pointSet(points []Point) map[Point]struct{} {
	set := map[Point]struct{}{}
	for _, p := range points {
		set[p] = struct{}{}
	}
	return set
}

func parseInput(input string) map[Point]int {
	var start, end Point
	grid := map[Point]int{}
	for r, row := range strings.Split(input, "\n") {
		for c, char := range row {
			if char == 'S' {
				start = Point{r, c}
			}
			if char == 'E' {
				end = Point{r, c}
			}
			if char != '#' {
				grid[Point{r, c}] = -1
			}
		}
	}
	grid[start] = 0
	curr := start
	for {
		if curr == end {
			break
		}
		for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := add(curr, dir)
			if _, ok := grid[next]; ok && grid[next] == -1 {
				grid[next] = grid[curr] + 1
				curr = next
			}
		}
	}
	return grid
}

func part1(input string, save int) int {
	grid := parseInput(input)
	cheats := []int{}

	for pos := range grid {
		for _, dir := range []Point{{0, 2}, {0, -2}, {2, 0}, {-2, 0}} {
			next := add(pos, dir)
			if _, ok := grid[next]; !ok {
				continue
			}
			if grid[pos]-grid[next] >= save+2 {
				cheats = append(cheats, grid[pos]-grid[next])
			}
		}
	}

	return len(cheats)
}

func part2(input string, save int) int {
	grid := parseInput(input)
	cheats := []int{}
	for pos := range grid {
		for radius := 2; radius < 21; radius++ {
			for dr := 0; dr < radius+1; dr++ {
				dc := radius - dr
				for dir := range pointSet([]Point{{dr, dc}, {dr, -dc}, {-dr, dc}, {-dr, -dc}}) {
					next := add(pos, dir)
					if _, ok := grid[next]; !ok {
						continue
					}
					if grid[pos]-grid[next] >= save+radius {
						cheats = append(cheats, grid[pos]-grid[next])
					}
				}
			}
		}
	}
	return len(cheats)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := string(input)
	part1_ans := part1(lines, 100)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines, 100)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
