package main

import (
	"fmt"
	"os"
	"strings"
	"utils"
)

func parseInput(input string) ([][]int, [][]int) {
	locks := [][]int{}
	keys := [][]int{}
	for _, line := range strings.Split(input, "\n\n") {
		count := utils.Repeat(-1, 5)
		for _, row := range strings.Split(line, "\n") {
			for j, char := range row {
				if char == '#' {
					count[j]++
				}
			}
		}
		if line[0] == '#' {
			locks = append(locks, count)
		} else {
			keys = append(keys, count)
		}
	}

	return locks, keys
}

func part1(input string) int {
	locks, keys := parseInput(input)
	fit := func(lock, key []int) bool {
		for i := range lock {
			if lock[i]+key[i] > 5 {
				return false
			}
		}
		return true
	}

	return utils.ApplySum(locks, func(lock []int) int {
		return utils.ApplySum(keys, func(key []int) int {
			if fit(lock, key) {
				return 1
			}
			return 0
		})
	})
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
}
