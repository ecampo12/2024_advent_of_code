package main

import (
	"fmt"
	"os"
	"strings"
	"utils"
)

type Point[T float32 | int] struct { // using generics to allow for both int and float32
	x T
	y T
}
type PointInt Point[int] // using a type alias to make declarations less verbose
type PointFloat Point[float32]

// Point functions to make it easier to read the code
func (p PointInt) IntToFloat() PointFloat {
	return PointFloat{float32(p.x), float32(p.y)}
}

func (p PointInt) AddInt(p2 PointInt) PointInt {
	return PointInt{p.x + p2.x, p.y + p2.y}
}

func (p PointFloat) AddFloat(p2 PointFloat) PointFloat {
	return PointFloat{p.x + p2.x, p.y + p2.y}
}

func (p PointInt) AddFloat(p2 PointFloat) PointFloat {
	return PointFloat{float32(p.x) + p2.x, float32(p.y) + p2.y}
}

func floodFill(grid map[PointInt]rune, current PointInt, letter rune, seen map[PointInt]bool) ([]PointInt, int) {
	if _, ok := grid[current]; !ok || grid[current] != letter {
		return []PointInt{}, 1
	}
	if _, ok := seen[current]; ok {
		return []PointInt{}, 0
	}
	seen[current] = true
	points := []PointInt{current}
	perimeter := 0

	for _, dir := range []PointInt{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		next := current.AddInt(dir)
		pts, p := floodFill(grid, next, letter, seen)
		points = append(points, pts...)
		perimeter += p
	}
	return points, perimeter
}

func part1(input []string) int {
	grid := map[PointInt]rune{}
	for r, rows := range input {
		for c, cols := range rows {
			grid[PointInt{r, c}] = cols
		}
	}

	seen := map[PointInt]bool{}
	prices := []int{}
	for point := range grid {
		if _, ok := seen[point]; ok {
			continue
		}
		area, perimeter := floodFill(grid, point, grid[point], seen)
		prices = append(prices, len(area)*perimeter)
	}

	return utils.Sum(prices)
}

func findSides(region map[PointFloat]bool) int {
	possible := map[PointFloat]bool{}
	corners := 0
	for point := range region {
		for _, dir := range []PointFloat{{-0.5, -0.5}, {0.5, -0.5}, {0.5, 0.5}, {-0.5, 0.5}} {
			possible[point.AddFloat(dir)] = true
		}
	}

	for point := range possible {
		locations := make([]int, 4)
		for d, dir := range []PointFloat{{-0.5, -0.5}, {0.5, -0.5}, {0.5, 0.5}, {-0.5, 0.5}} {
			if _, ok := region[point.AddFloat(dir)]; ok {
				locations[d] = 1
			} else {
				locations[d] = 0
			}
		}
		num := utils.Sum(locations)
		if num == 1 || num == 3 {
			corners++
		}
		if num == 2 {
			if locations[0] == locations[2] || locations[1] == locations[3] {
				corners += 2
			}
		}
	}
	return corners
}

func part2(input []string) int {
	grid := map[PointInt]rune{}
	for r, rows := range input {
		for c, cols := range rows {
			grid[PointInt{r, c}] = cols
		}
	}

	seen := map[PointInt]bool{}
	prices := []int{}
	for point := range grid {
		if _, ok := seen[point]; ok {
			continue
		}
		area, _ := floodFill(grid, point, grid[point], seen)
		region := map[PointFloat]bool{}
		for _, p := range area {
			region[p.IntToFloat()] = true
		}
		prices = append(prices, len(area)*findSides(region))
	}

	return utils.Sum(prices)
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
