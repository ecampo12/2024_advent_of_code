package main

import (
	"fmt"
	"os"
	"strings"
	"utils"
)

type Point struct{ x, y int }

func add(a, b Point) Point { return Point{a.x + b.x, a.y + b.y} }
func sub(a, b Point) Point { return Point{a.x - b.x, a.y - b.y} }

var NUMERIC_KEYPAD = map[rune]Point{
	'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
	'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
	'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
	'0': {3, 1}, 'A': {3, 2},
}

var DIRECTIONAL_KEYPAD = map[rune]Point{
	'^': {0, 1}, 'A': {0, 2},
	'<': {1, 0}, 'v': {1, 1}, '>': {1, 2},
}

var MOVES = map[rune]Point{'^': {-1, 0}, 'v': {1, 0}, '<': {0, -1}, '>': {0, 1}}

func permuteMoveset(moveset []rune) []string {
	if len(moveset) == 0 {
		return []string{""}
	}

	moves := map[string]bool{}
	for i, m := range moveset {
		rest := append([]rune{}, moveset[:i]...)
		rest = append(rest, moveset[i+1:]...)
		for _, p := range permuteMoveset(rest) {
			moves[string(m)+p] = true
		}
	}
	list := []string{}
	for k := range moves {
		list = append(list, k)
	}
	return list
}

func sequenceToMoveset(start, end Point, keypad map[rune]Point) []string {
	moves := []rune{}
	d := sub(end, start)

	if d.x < 0 {
		moves = append(moves, utils.Repeat('^', d.x)...)
	} else {
		moves = append(moves, utils.Repeat('v', d.x)...)
	}

	if d.y < 0 {
		moves = append(moves, utils.Repeat('<', d.y)...)
	} else {
		moves = append(moves, utils.Repeat('>', d.y)...)
	}

	validSeq := []string{}
	for _, p := range permuteMoveset(moves) {
		pos := []Point{start}
		valid := true
		for _, m := range p {
			next := add(pos[len(pos)-1], MOVES[m])
			pos = pos[:len(pos)-1]
			if !utils.MapContainsValue(keypad, next) {
				valid = false
				break
			}
			pos = append(pos, next)
		}
		if valid {
			validSeq = append(validSeq, p+"A")
		}
	}
	return validSeq
}

type Node struct {
	seq   string
	limit int
	depth int
}

var cache = map[Node]int{}

func minLength(seq string, limit int, depth int) int {
	if _, ok := cache[Node{seq, limit, depth}]; ok {
		return cache[Node{seq, limit, depth}]
	}

	var keypad map[rune]Point
	keypad = DIRECTIONAL_KEYPAD
	if depth == 0 {
		keypad = NUMERIC_KEYPAD
	}
	current := keypad['A']

	length := []int{}
	for _, char := range seq {
		next := keypad[char]
		movesets := sequenceToMoveset(current, next, keypad)
		if depth >= limit {
			length = append(length, len(utils.SliceMin(movesets, func(s string) int { return len(s) })))
		} else {
			values := utils.Apply(movesets, func(s string) int { return minLength(s, limit, depth+1) })
			length = append(length, utils.SliceMin(values, func(s int) int { return s }))
		}
		current = next
	}
	cache[Node{seq, limit, depth}] = utils.Sum(length)
	return utils.Sum(length)
}

func part1(input []string) int {
	return utils.ApplySum(input, func(s string) int { return minLength(s, 2, 0) * utils.Int(s[:len(s)-1]) })
}

func part2(input []string) int {
	return utils.ApplySum(input, func(s string) int { return minLength(s, 25, 0) * utils.Int(s[:len(s)-1]) })
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
