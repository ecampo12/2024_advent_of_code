package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"utils"
)

func part1(input []string) int {
	a := make([]int, len(input))
	b := make([]int, len(input))

	for i, line := range input {
		fmt.Sscanf(line, "%d %d", &a[i], &b[i])
	}

	sort.Ints(a)
	sort.Ints(b)

	sum := 0
	for i := range a {
		sum += utils.Abs(a[i] - b[i])
	}
	return sum
}

func count(a []int) map[int]int {
	m := make(map[int]int)
	for _, n := range a {
		m[n]++
	}
	return m
}

func part2(input []string) int {
	a := make([]int, len(input))
	b := make([]int, len(input))

	for i, line := range input {
		fmt.Sscanf(line, "%d %d", &a[i], &b[i])
	}

	c := utils.Counter(b)
	sum := 0
	for _, n := range a {
		sum += n * c[n]
	}

	return sum
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
