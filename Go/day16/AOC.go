package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	row, col int
}

func add(p1 Point, p2 Point) Point {
	return Point{p1.row + p2.row, p1.col + p2.col}
}

type Item struct {
	value     Point
	cost      int
	index     int
	direction Point
	path      []Point
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func parseInput(input string) ([]string, Point, Point) {
	grid := strings.Split(string(input), "\n")
	var start Point
	var end Point
	for r, row := range grid {
		for c, char := range row {
			if char == 'S' {
				start = Point{r, c}
			} else if char == 'E' {
				end = Point{r, c}
			}
		}
	}
	return grid, start, end
}

type Node struct {
	point Point
	d     Point
}

func dijkstra(grid []string, start Point, end Point) int {
	rows, cols := len(grid), len(grid[0])
	pq := PriorityQueue{}
	seen := make(map[Node]bool)
	pq.Push(&Item{value: start, cost: 0, direction: Point{0, 1}})

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item)
		if curr.value == end {
			return curr.cost
		}

		if seen[Node{curr.value, curr.direction}] {
			continue
		}
		seen[Node{curr.value, curr.direction}] = true

		d := add(curr.value, curr.direction)
		nr, nc := d.row, d.col
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] != '#' {
			heap.Push(&pq, &Item{value: Point{nr, nc}, cost: curr.cost + 1, direction: curr.direction})
		}

		heap.Push(&pq, &Item{value: curr.value, cost: curr.cost + 1000, direction: Point{-curr.direction.col, curr.direction.row}})
		heap.Push(&pq, &Item{value: curr.value, cost: curr.cost + 1000, direction: Point{curr.direction.col, -curr.direction.row}})
	}
	return -1
}

func part1(input string) int {
	return dijkstra(parseInput(input))
}

func findAllShortestPaths(grid []string, start Point, end Point) map[Point]bool {
	rows, cols := len(grid), len(grid[0])
	pq := PriorityQueue{}
	seen := make(map[Node]bool)
	pq.Push(&Item{value: start, cost: 0, direction: Point{0, 1}, path: []Point{start}})
	points := make(map[Point]bool)
	best := math.MaxInt64

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item)
		seen[Node{curr.value, curr.direction}] = true

		if curr.value == end {
			if curr.cost <= best {
				best = curr.cost
				for _, p := range curr.path {
					points[p] = true
				}
			} else {
				break
			}
		}

		d := add(curr.value, curr.direction)
		nr, nc := d.row, d.col
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] != '#' && !seen[Node{Point{nr, nc}, curr.direction}] {
			heap.Push(&pq, &Item{value: Point{nr, nc}, cost: curr.cost + 1, direction: curr.direction, path: append(curr.path, Point{nr, nc})})
		}

		dr, dc := curr.direction.row, curr.direction.col
		if !seen[Node{curr.value, Point{-dc, dr}}] {
			heap.Push(&pq, &Item{value: curr.value, cost: curr.cost + 1000, direction: Point{-dc, dr}, path: append([]Point{}, curr.path...)})
		}
		if !seen[Node{curr.value, Point{dc, -dr}}] {
			heap.Push(&pq, &Item{value: curr.value, cost: curr.cost + 1000, direction: Point{dc, -dr}, path: append([]Point{}, curr.path...)})
		}
	}
	return points
}

func part2(input string) int {
	return len(findAllShortestPaths(parseInput(input)))
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
