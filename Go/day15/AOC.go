package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	r, c int
}

type Points []Point

func (p Points) contains(point Point) bool {
	for _, p := range p {
		if p == point {
			return true
		}
	}
	return false
}

type Grid map[Point]rune

func parseInput(input string) (Grid, []rune, Point) {
	grid := make(Grid, 0)
	var start Point
	var moves []rune
	parts := strings.Split(input, "\n\n")
	for r, line := range strings.Split(parts[0], "\n") {
		for c, char := range line {
			grid[Point{r, c}] = char
			if char == '@' {
				start = Point{r, c}
			}
		}
	}
	moves = []rune(strings.ReplaceAll(parts[1], "\n", ""))
	return grid, moves, start
}

func (g Grid) calculateGPS() int {
	sum := 0
	for point, char := range g {
		if char == 'O' || char == '[' {
			sum += 100*point.r + point.c
		}
	}
	return sum
}

func part1(input string) int {
	grid, moves, start := parseInput(input)
	directions := map[rune]Point{
		'^': {-1, 0},
		'v': {1, 0},
		'<': {0, -1},
		'>': {0, 1},
	}

	for _, move := range moves {
		boxes := Points{}
		dr, dc := directions[move].r, directions[move].c
		cr, cc := start.r, start.c
		moveable := true
		for true {
			cr += dr
			cc += dc
			char := grid[Point{cr, cc}]
			if char == '#' {
				moveable = false
				break
			}
			if char == '.' {
				break
			}
			if char == 'O' {
				boxes = append(boxes, Point{cr, cc})
			}
		}
		if !moveable {
			continue
		}
		grid[start] = '.'
		grid[Point{start.r + dr, start.c + dc}] = '@'
		for _, box := range boxes {
			grid[Point{box.r + dr, box.c + dc}] = 'O'
		}
		start = Point{start.r + dr, start.c + dc}
	}

	return grid.calculateGPS()
}

func expandGrid(input string) string {
	sb := strings.Builder{}
	expansion := map[rune]string{'.': "..", '#': "##", 'O': "[]", '@': "@."}
	for _, line := range strings.Split(input, "\n") {
		for _, c := range line {
			sb.WriteString(expansion[c])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	expanded := expandGrid(parts[0])
	grid, moves, start := parseInput(expanded + "\n\n" + parts[1])
	directions := map[rune]Point{
		'^': {-1, 0},
		'v': {1, 0},
		'<': {0, -1},
		'>': {0, 1},
	}
	for _, move := range moves {
		boxes := Points{start}
		dr, dc := directions[move].r, directions[move].c
		moveable := true
		for i := 0; i < len(boxes); i++ {
			nr := boxes[i].r + dr
			nc := boxes[i].c + dc
			if boxes.contains(Point{nr, nc}) {
				continue
			}
			char := grid[Point{nr, nc}]
			if char == '#' {
				moveable = false
				break
			}
			if char == '[' {
				boxes = append(boxes, Point{nr, nc})
				boxes = append(boxes, Point{nr, nc + 1})
			}
			if char == ']' {
				boxes = append(boxes, Point{nr, nc})
				boxes = append(boxes, Point{nr, nc - 1})
			}
		}
		if !moveable {
			continue
		}
		chars := map[Point]rune{}
		grid[start] = '.'
		for _, box := range boxes[1:] {
			chars[box] = grid[box]
			grid[box] = '.'
		}

		for _, box := range boxes[1:] {
			grid[Point{box.r + dr, box.c + dc}] = chars[box]
		}
		grid[Point{start.r + dr, start.c + dc}] = '@'
		start = Point{start.r + dr, start.c + dc}
	}

	return grid.calculateGPS()

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
