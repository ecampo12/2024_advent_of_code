package main

import (
	"fmt"
	"os"
	"strings"
)

var cache = map[string]int{}

func countCombos(towel string, patterns []string) int {
	if _, ok := cache[towel]; ok {
		return cache[towel]
	}
	if len(towel) == 0 {
		return 1
	}
	count := 0
	for _, p := range patterns {
		if strings.HasPrefix(towel, p) {
			count += countCombos(towel[len(p):], patterns)
		}
	}
	cache[towel] = count
	return count
}

func part1(input string) int {
	towelPatterns := strings.Split(strings.Split(input, "\n\n")[0], ", ")
	desiredPatters := strings.Split(strings.Split(input, "\n\n")[1], "\n")
	count := 0
	for _, t := range desiredPatters {
		c := countCombos(t, towelPatterns)
		if c > 0 {
			count++
		}
	}

	return count
}

func part2(input string) int {
	towelPatterns := strings.Split(strings.Split(input, "\n\n")[0], ", ")
	desiredPatters := strings.Split(strings.Split(input, "\n\n")[1], "\n")
	count := 0
	for _, t := range desiredPatters {
		count += countCombos(t, towelPatterns)
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
	part1_ans := part1(lines)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
