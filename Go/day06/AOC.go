package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"utils"
)

type Point struct {
	r, c   int
	dr, dc int
}

type Grid struct {
	grid          [][]rune
	width, height int
	start         Point
}

func (g *Grid) isInBounds(p Point) bool {
	return p.r >= 0 && p.r < g.height && p.c >= 0 && p.c < g.width
}

func parseInput(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	start := Point{0, 0, 0, 0}
	for i, line := range lines {
		if strings.ContainsRune(line, '^') {
			start = Point{i, strings.IndexRune(line, '^'), -1, 0}
		}
		grid[i] = []rune(line)
	}
	return Grid{grid, len(grid[0]), len(grid), start}
}

func (g *Grid) traverse() map[Point]bool {
	dr, dc := g.start.dr, g.start.dc
	seen := map[Point]bool{}
	r, c := g.start.r, g.start.c
	for true {
		seen[Point{r, c, 0, 0}] = true
		if !g.isInBounds(Point{r + dr, c + dc, 0, 0}) {
			break
		}

		if g.grid[r+dr][c+dc] == '#' {
			dc, dr = -dr, dc
		} else {
			r += dr
			c += dc
		}
	}
	return seen
}

func (g *Grid) findLoop(obstacle Point) bool {
	dr, dc := -1, 0
	seen := map[Point]bool{}
	r, c := g.start.r, g.start.c
	for true {

		seen[Point{r, c, dr, dc}] = true
		if !g.isInBounds(Point{r + dr, c + dc, 0, 0}) {
			return false
		}

		if g.grid[r+dr][c+dc] == '#' || (r+dr == obstacle.r && c+dc == obstacle.c) {
			dc, dr = -dr, dc
		} else {
			r += dr
			c += dc
		}
		if seen[Point{r, c, dr, dc}] {
			return true
		}

	}
	return false
}

func part1(input string) int {
	grid := parseInput(input)
	return len(grid.traverse())
}

func part2(input string) int {
	grid := parseInput(input)
	paths := grid.traverse()
	count := 0
	for p := range paths {
		if grid.findLoop(p) {
			count++
		}
	}
	return count
}

func part2WithGorountines(input string) int {
	grid := parseInput(input)
	paths := grid.traverse()

	results := make(chan int)
	var wg sync.WaitGroup

	for p := range paths {
		wg.Add(1)
		go func(path Point) {
			defer wg.Done()
			if grid.findLoop(path) {
				results <- 1
			} else {
				results <- 0
			}
		}(p)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	count := 0
	for res := range results {
		count += res
	}
	return count
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := string(input)
	fmt.Printf("Part 1: %d\n", part1(lines))
	utils.Timer(func() {
		fmt.Printf("Part 2: %d\n", part2(lines))
	}, "Part 2")
	fmt.Println()
	utils.Timer(func() {
		fmt.Printf("Part 2 with Goroutines: %d\n", part2WithGorountines(lines))
	}, "Part 2 with Goroutines")
}
