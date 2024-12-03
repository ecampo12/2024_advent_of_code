package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input []string) [][]int {
	reports := make([][]int, len(input))
	for i, line := range input {
		nums := strings.Fields(line)
		reports[i] = make([]int, len(nums))
		for j, num := range nums {
			fmt.Sscanf(num, "%d", &reports[i][j])
		}
	}

	return reports
}

func isSafe(report []int) bool {
	increasing := report[0] < report[1]
	decreasing := report[0] > report[1]
	acceptable := map[int]bool{1: true, 2: true, 3: true, -1: true, -2: true, -3: true}

	for i := range len(report) - 1 {
		diff := report[i+1] - report[i]
		if !acceptable[diff] || diff == 0 || (increasing && diff < 0) || (decreasing && diff > 0) {
			return false
		}
	}

	return true
}

func part1(input []string) int {
	reports := parseInput(input)
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	return count
}

func part2(input []string) int {
	reports := parseInput(input)
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		} else {
			for i := 0; i < len(report); i++ {
				r := append([]int{}, report...) // Copy the slice to avoid modifying the original
				r = append(r[:i], r[i+1:]...)
				if isSafe(r) {
					count++
					break
				}
			}
		}
	}
	return count
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
