package main

import (
	"fmt"
	"os"
	"strings"
	"utils"
)

type point struct {
	x, y int
}

func (p point) add(p2 point) point {
	return point{p.x + p2.x, p.y + p2.y}
}

func trailheadScore(hikingTrail map[point]rune, trailhead point) int {
	summits := 0
	seen := make(map[point]bool)
	directions := []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := []point{trailhead}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			next := current.add(dir)
			if _, ok := hikingTrail[next]; !ok {
				continue
			}
			if hikingTrail[next] != hikingTrail[current]+1 {
				continue
			}

			if _, ok := seen[next]; ok {
				continue
			}

			seen[next] = true
			if hikingTrail[next] == '9' {
				summits++
			} else {
				queue = append(queue, next)
			}
		}
	}
	return summits
}

func part1(input []string) int {
	hikingTrail := make(map[point]rune)
	trailheads := []point{}
	for r, row := range input {
		for c, col := range row {
			hikingTrail[point{r, c}] = col
			if col == '0' {
				trailheads = append(trailheads, point{r, c})
			}
		}
	}
	scores := []int{}
	for _, trailhead := range trailheads {
		scores = append(scores, trailheadScore(hikingTrail, trailhead))
	}

	return utils.Sum(scores)
}

func trailheadRating(hikingTrail map[point]rune, trailhead point) int {
	trails := 0
	seen := map[point]int{trailhead: 1}
	directions := []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := []point{trailhead}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if hikingTrail[current] == '9' {
			trails += seen[current]
		}
		for _, dir := range directions {
			next := current.add(dir)
			if _, ok := hikingTrail[next]; !ok {
				continue
			}
			if hikingTrail[next] != hikingTrail[current]+1 {
				continue
			}
			if _, ok := seen[next]; ok {
				seen[next] += seen[current]
				continue
			}
			seen[next] = seen[current]
			queue = append(queue, next)
		}
	}
	return trails
}

func part2(input []string) int {
	hikingTrail := make(map[point]rune)
	trailheads := []point{}
	for r, row := range input {
		for c, col := range row {
			hikingTrail[point{r, c}] = col
			if col == '0' {
				trailheads = append(trailheads, point{r, c})
			}
		}
	}
	scores := []int{}
	for _, trailhead := range trailheads {
		scores = append(scores, trailheadRating(hikingTrail, trailhead))
	}

	return utils.Sum(scores)
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
