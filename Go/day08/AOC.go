package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func (p *Point) update(x, y int) {
	p.x = x + p.x
	p.y = y + p.y
}

type Map struct {
	width, height int
	antennas      map[rune][]Point
}

func (m *Map) inBounds(p Point) bool {
	return p.x >= 0 && p.x < m.width && p.y >= 0 && p.y < m.height
}

func ParseMap(input []string) Map {
	m := Map{width: len(input[0]), height: len(input), antennas: make(map[rune][]Point)}
	for r, line := range input {
		for c, char := range line {
			if char != '.' {
				m.antennas[char] = append(m.antennas[char], Point{r, c})
			}
		}
	}
	return m
}

// Uses channels to generate all pairs of points
func generatePoints(p []Point) chan []Point {
	c := make(chan []Point)
	go func() {
		defer close(c)
		for i := 0; i < len(p); i++ {
			for j := i + 1; j < len(p); j++ {
				c <- []Point{p[i], p[j]}
			}
		}
	}()
	return c
}

func part1(input []string) int {
	m := ParseMap(input)
	antiNodes := make(map[Point]bool)
	for _, v := range m.antennas {
		for pair := range generatePoints(v) {
			antiNodes[Point{2*pair[1].x - pair[0].x, 2*pair[1].y - pair[0].y}] = true
			antiNodes[Point{2*pair[0].x - pair[1].x, 2*pair[0].y - pair[1].y}] = true
		}
	}
	for k := range antiNodes {
		if !m.inBounds(k) {
			delete(antiNodes, k)
		}
	}
	return len(antiNodes)
}

func (m *Map) pointsOnLine(p1, p2 Point) map[Point]bool {
	points := make(map[Point]bool)
	dr := p2.x - p1.x
	dc := p2.y - p1.y
	for m.inBounds(p1) {
		points[p1] = true
		p1.update(dr, dc)
	}
	return points
}

func part2(input []string) int {
	m := ParseMap(input)
	antiNodes := make(map[Point]bool)
	for _, v := range m.antennas {
		for pair := range generatePoints(v) {
			for k := range m.pointsOnLine(pair[0], pair[1]) {
				antiNodes[k] = true
			}
			for k := range m.pointsOnLine(pair[1], pair[0]) {
				antiNodes[k] = true
			}
		}
	}

	return len(antiNodes)
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
