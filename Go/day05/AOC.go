package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	a string
	b string
}

type Update struct {
	pages []string
	rules []Rule
}

func (u *Update) isValid() bool {
	indexUpdate := map[string]int{}
	for i, p := range u.pages {
		indexUpdate[p] = i
	}

	for _, rule := range u.rules {
		if indexUpdate[rule.a] > indexUpdate[rule.b] {
			return false
		}
	}
	return true
}

// Topological sort of the pages using the rules and using the Kahn's algorithm.
// I checked, there is no cycle in the graph, so the topological sort is unique.
// https://en.wikipedia.org/wiki/Topological_sorting#Kahn's_algorithm
func (u *Update) topologicalSort() []string {
	graph := map[string][]string{}
	inDegree := map[string]int{}

	nodes := map[string]bool{}
	for _, rule := range u.rules {
		if _, ok := graph[rule.a]; !ok {
			graph[rule.a] = []string{}
		}
		graph[rule.a] = append(graph[rule.a], rule.b)
		inDegree[rule.b]++
		nodes[rule.a] = true
		nodes[rule.b] = true
	}

	for node := range nodes {
		if _, ok := inDegree[node]; !ok {
			inDegree[node] = 0
		}
	}

	queue := []string{}
	sorted := []string{}
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sorted = append(sorted, node)
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return sorted
}

func parseInput(input string) []Update {
	updates := []Update{}
	s := strings.Split(input, "\n\n")
	rules := []Rule{}
	for _, rule := range strings.Split(s[0], "\n") {
		p := strings.Split(rule, "|")
		rules = append(rules, Rule{p[0], p[1]})
	}
	for _, line := range strings.Split(s[1], "\n") {

		update := []string{}
		set := map[string]bool{}
		for _, u := range strings.Split(line, ",") {
			update = append(update, u)
			set[u] = true
		}
		// Find the rules that are applicable to this update
		localRules := []Rule{}
		for _, rule := range rules {
			if set[rule.a] && set[rule.b] {
				localRules = append(localRules, rule)
			}
		}
		updates = append(updates, Update{update, localRules})
	}
	return updates
}

func part1(input string) int {
	updates := parseInput(input)
	sum := 0
	for _, update := range updates {
		if update.isValid() { // we can use the topological sort to check if the update is valid, not sure if it is faster
			mid, _ := strconv.Atoi(update.pages[len(update.pages)/2])
			sum += mid
		}
	}
	return sum
}

func part2(input string) int {
	updates := parseInput(input)
	sum := 0
	for _, update := range updates {
		if !update.isValid() {
			mid, _ := strconv.Atoi(update.topologicalSort()[len(update.pages)/2])
			sum += mid
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
