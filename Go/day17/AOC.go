package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func toString(nums []int) string {
	strs := []string{}
	for _, n := range nums {
		strs = append(strs, strconv.Itoa(n))
	}
	return strings.Join(strs, ",")
}

func parseInput(input string) []int {
	nums := []int{}
	for _, n := range regexp.MustCompile(`\d+`).FindAllString(input, -1) {
		x, _ := strconv.Atoi(n)
		nums = append(nums, x)
	}
	return nums
}

func run(registers []int, program []int) string {
	A := registers[0]
	B := registers[1]
	C := registers[2]
	combo := func(val int) int {
		if val <= 3 {
			return val
		}
		switch val {
		case 4:
			return A
		case 5:
			return B
		case 6:
			return C
		}
		return 0
	}

	output := []string{}
	ip := 0
	for ip < len(program) {
		opcode := program[ip]
		operand := program[ip+1]
		switch opcode {
		case 0:
			A = A >> combo(operand)
		case 1:
			B = B ^ operand
		case 2:
			B = combo(operand) % 8
		case 3:
			if A != 0 {
				ip = operand
				continue
			}
		case 4:
			B = B ^ C
		case 5:
			output = append(output, strconv.Itoa(combo(operand)%8))
		case 6:
			B = A >> combo(operand)
		case 7:
			C = A >> combo(operand)
		}
		ip += 2
	}

	return strings.Join(output, ",")
}

func part1(input string) string {
	nums := parseInput(input)
	return run(nums[:3], nums[3:])
}

// Apprently what we are looking for is a quine. A quine is a program that outputs its own source code.
// https://en.wikipedia.org/wiki/Quine_(computing)
func findQuine(registers []int, program []int) int {
	candidates := []int{0}
	for i := len(program) - 1; i >= 0; i-- {
		nextcandidates := []int{}
		for _, val := range candidates {
			for j := 0; j < 8; j++ {
				registers[0] = (val << 3) + j
				if run(registers, program) == toString(program[i:]) {
					nextcandidates = append(nextcandidates, registers[0])
				}
			}
			candidates = nextcandidates
		}
	}
	return candidates[0]
}

func part2(input string) int {
	nums := parseInput(input)
	return findQuine(nums[:3], nums[3:])
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := string(input)
	part1_ans := part1(lines)
	fmt.Printf("Part 1: %s\n", part1_ans)

	part2_ans := part2(lines)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
