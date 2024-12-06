package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) (map[int][]int, [][]int) {
	s := strings.Split(input, "\n\n")
	rules := map[int][]int{}
	for _, rule := range strings.Split(s[0], "\n") {
		p := strings.Split(rule, "|")
		a, _ := strconv.Atoi(p[0])
		b, _ := strconv.Atoi(p[1])
		if rules[a] == nil {
			rules[a] = []int{}
		}
		rules[a] = append(rules[a], b)
	}
	updates := make([][]int, len(strings.Split(s[1], "\n")))
	for i, line := range strings.Split(s[1], "\n") {
		update := []int{}
		for _, u := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(u)
			update = append(update, num)
		}
		updates[i] = update
	}
	return rules, updates
}
func valid(rules map[int][]int, update []int) bool {
	updateIndex := map[int]int{}
	for i, u := range update {
		updateIndex[u] = i
	}

	// for _, u := range update {

	// }
	return true
}

func part1(input string) int {
	rules, updates := parseInput(input)
	sum := 0
	for _, update := range updates {
		localRules := make(map[int][]int)
		for _, r := range update {
			localRules[r] = rules[r]
		}
		if valid(localRules, update) {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func part2(input string) int {
	return 0
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
