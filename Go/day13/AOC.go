package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"utils"
)

type Machine struct {
	ax, ay int
	bx, by int
	px, py int
}

func parseInput(input string) []Machine {
	parts := strings.Split(input, "\n\n")
	machines := []Machine{}
	for _, part := range parts {
		re := regexp.MustCompile(`(\d+)`)
		matches := re.FindAllString(part, -1)
		nums := make([]int, len(matches))
		for i, match := range matches {
			fmt.Sscanf(match, "%d", &nums[i])
		}

		ax, ay, bx, by, px, py := nums[0], nums[1], nums[2], nums[3], nums[4], nums[5]
		machines = append(machines, Machine{ax, ay, bx, by, px, py})
	}

	return machines
}

// Cramer's Rule: https://www.cuemath.com/algebra/cramers-rule/
func CramersRule(A [2][2]int, B [2]int) (int, int) {
	Det := A[0][0]*A[1][1] - A[0][1]*A[1][0]
	DetX := B[0]*A[1][1] - B[1]*A[0][1]
	DetY := A[0][0]*B[1] - A[1][0]*B[0]

	return DetX / Det, DetY / Det
}

func part1(input string) int {
	machines := parseInput(input)
	tokens := []int{}
	for _, m := range machines {
		A := [2][2]int{{m.ax, m.bx}, {m.ay, m.by}}
		B := [2]int{m.px, m.py}
		x, y := CramersRule(A, B)
		if m.ax*x+m.bx*y == m.px && m.ay*x+m.by*y == m.py {
			tokens = append(tokens, 3*x+y)
		}
	}
	return utils.Sum(tokens)
}

func part2(input string) int {
	machines := parseInput(input)
	tokens := []int{}
	for _, m := range machines {
		m.px += 10_000_000_000_000
		m.py += 10_000_000_000_000
		A := [2][2]int{{m.ax, m.bx}, {m.ay, m.by}}
		B := [2]int{m.px, m.py}
		x, y := CramersRule(A, B)
		if m.ax*x+m.bx*y == m.px && m.ay*x+m.by*y == m.py {
			tokens = append(tokens, 3*x+y)
		}
	}
	return utils.Sum(tokens)
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
