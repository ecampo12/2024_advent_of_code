package main

import (
	"fmt"
	"os"
	"strings"
	"utils"
)

type Point struct {
	row, col int
}

func add(a, b Point) Point {
	return Point{a.row + b.row, a.col + b.col}
}

type Path struct {
	point Point
	steps int
}

func parseInput(input string) []Point {
	points := []Point{}
	for _, line := range strings.Split(input, "\n") {
		var p Point
		fmt.Sscanf(line, "%d,%d", &p.row, &p.col)
		points = append(points, p)
	}
	return points
}

func run(points []Point, bytes, size int) int {
	points = points[:bytes]
	grid := map[Point]rune{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[Point{i, j}] = '.'
		}
	}

	for _, p := range points {
		grid[p] = '#'
	}

	exit := Point{size - 1, size - 1}
	seen := map[Point]bool{}
	seen[Point{0, 0}] = true
	queue := []Path{{Point{0, 0}, 0}}
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.point == exit {
			return current.steps
		}

		for _, dir := range directions {
			next := add(current.point, dir)
			if _, ok := grid[next]; ok && grid[next] == '.' && !seen[next] {
				seen[next] = true
				queue = append(queue, Path{next, current.steps + 1})
			}
		}
	}
	return -1
}

func part1(input string, bytes, size int) int {
	return run(parseInput(input), bytes, size)
}

func part2(input string, size int) string {
	points := parseInput(input)
	lo, hi := 0, len(points)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if run(points, mid, size) == -1 {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return fmt.Sprintf("%d,%d", points[lo-1].row, points[lo-1].col)
}

func run2(grid map[Point]rune, exit Point) int {
	seen := map[Point]bool{}
	seen[Point{0, 0}] = true
	queue := []Path{{Point{0, 0}, 0}}
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.point == exit {
			return current.steps
		}

		for _, dir := range directions {
			next := add(current.point, dir)
			if _, ok := grid[next]; ok && grid[next] == '.' && !seen[next] {
				seen[next] = true
				queue = append(queue, Path{next, current.steps + 1})
			}
		}
	}
	return -1
}

func part2Alt(input string, size int) string {
	points := parseInput(input)
	grid := map[Point]rune{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[Point{i, j}] = '.'
		}
	}

	for _, p := range points {
		grid[p] = '#'
	}

	exit := Point{size - 1, size - 1}

	for i := len(points) - 1; i >= 0; i-- {
		grid[points[i]] = '.'
		if run2(grid, exit) != -1 {
			return fmt.Sprintf("%d,%d", points[i].row, points[i].col)
		}
	}

	return ""
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := string(input)
	part1_ans := part1(lines, 1024, 71)
	fmt.Printf("Part 1: %d\n", part1_ans)

	utils.Timer(func() {
		part2_ans := part2(lines, 71)
		fmt.Printf("Part 2: %s\n", part2_ans)
	}, "Part 2")
	fmt.Print("\n")
	utils.Timer(func() {
		part2Alt_ans := part2Alt(lines, 71)
		fmt.Printf("Part 2: %s\n", part2Alt_ans)
	}, "Part 2 Alt")
}
