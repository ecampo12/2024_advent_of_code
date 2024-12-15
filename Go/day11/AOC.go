package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntMap map[int]int

func (m IntMap) Add(n int, count int) {
	if _, ok := m[n]; !ok {
		m[n] = 0
	}
	m[n] += count
}

func (m IntMap) size() int {
	size := 0
	for _, count := range m {
		size += count
	}
	return size
}

// keep track of the unique stones that appear at each step
// means that we do less work when expanding the stones, and need less memory to store them
func expandStones(stones IntMap) IntMap {
	new_stones := IntMap{}
	for num, count := range stones {
		if num == 0 {
			new_stones.Add(1, count)
			continue
		}
		str := strconv.Itoa(num)
		if len(str)%2 == 0 {
			left, _ := strconv.Atoi(str[:len(str)/2])
			right, _ := strconv.Atoi(str[len(str)/2:])
			new_stones.Add(left, count)
			new_stones.Add(right, count)
		} else {
			new_stones.Add(num*2024, count)
		}
	}
	return new_stones
}

func part1(input string, blinks int) int {
	stones := IntMap{}
	for _, num := range strings.Fields(input) {
		n, _ := strconv.Atoi(num)
		stones.Add(n, 1)
	}

	for i := 0; i < blinks; i++ {
		stones = expandStones(stones)
	}

	return stones.size()
}

func part2(input string) int {
	stones := IntMap{}
	for _, num := range strings.Fields(input) {
		n, _ := strconv.Atoi(num)
		stones.Add(n, 1)
	}

	for i := 0; i < 75; i++ {
		stones = expandStones(stones)
	}

	return stones.size()
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := string(input)
	part1_ans := part1(lines, 25)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
