package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1(input string) int {
	sum := 0
	for _, c := range regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch(input, -1) {
		a, b := c[1], c[2]
		num1, _ := strconv.Atoi(a)
		num2, _ := strconv.Atoi(b)
		sum = num1 * num2
	}
	return sum
}

func part2(input string) int {
	enabled := true
	sum := 0
	for _, c := range regexp.MustCompile(`don't|do|mul\(\d+,\d+\)`).FindAllStringSubmatch(input, -1) {
		if c[0] == "do" {
			enabled = true
		} else if c[0] == "don't" {
			enabled = false
		} else if enabled {
			sum += part1(c[0])
		}
	}
	return sum
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
