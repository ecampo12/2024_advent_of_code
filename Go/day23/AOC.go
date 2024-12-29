package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"utils"
)

type Graph map[string][]string

func contains(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func buildGraph(input []string) Graph {
	connection := Graph{}
	for _, line := range input {
		parts := strings.Split(line, "-")
		if _, ok := connection[parts[0]]; !ok {
			connection[parts[0]] = []string{}
		}
		if _, ok := connection[parts[1]]; !ok {
			connection[parts[1]] = []string{}
		}
		connection[parts[0]] = append(connection[parts[0]], parts[1])
		connection[parts[1]] = append(connection[parts[1]], parts[0])
	}
	return connection
}

func part1(input []string) int {
	connection := buildGraph(input)
	triples := map[string][]string{}

	for k, v := range connection {
		for _, c := range v {
			for _, d := range connection[c] {
				if contains(v, d) {
					if k[0] == 't' || c[0] == 't' || d[0] == 't' {
						key := []string{k, c, d}
						sort.Strings(key)
						triples[strings.Join(key, "-")] = []string{k, c, d}
					}
				}
			}
		}
	}
	return len(triples)
}

var passwords = utils.NewSet[string]()

func buildSet(connections Graph, conn string, group utils.Set[string]) {
	key := group.ToSlice()
	sort.Strings(key)
	password := strings.Join(key, ",")
	if passwords[password] {
		return
	}
	passwords.Add(password)
	for _, neighbor := range connections[conn] {
		if group[neighbor] {
			continue
		}
		if !utils.All(group.ToSlice(), func(n string) bool { return contains(connections[neighbor], n) }) {
			continue
		}
		group.Add(neighbor)
		buildSet(connections, neighbor, group)
	}
}

func part2(input []string) string {
	connections := buildGraph(input)
	for k := range connections {
		group := utils.NewSet[string]()
		group.Add(k)
		buildSet(connections, k, group)
	}

	return utils.SliceMax(passwords.ToSlice(), func(s string) int { return len(s) })
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
	fmt.Printf("Part 2: %s\n", part2_ans)
}
