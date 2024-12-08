package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"utils"
)

func parseInput(input []string) [][]int {
	res := make([][]int, len(input))
	for i, line := range input {
		vals := regexp.MustCompile(`\d+`).FindAllString(line, -1)
		nums := make([]int, len(vals))
		for n, num := range vals {
			x, _ := strconv.Atoi(num)
			nums[n] = x
		}
		res[i] = nums
	}
	return res
}

// Brute force all possible combinations of operations using bit manipulation
func part1(input []string) int {
	equations := parseInput(input)
	sum := 0
	for _, eq := range equations {
		target := eq[0]
		eq = eq[1:]
		for n := range int(math.Pow(2, float64(len(eq)-1))) {
			res := eq[0]
			for i := 0; i < len(eq)-1; i++ {
				b := eq[i+1]
				op := byte((n >> i) & 0x1)
				switch op {
				case 0x0:
					res += b
				case 0x1:
					res *= b
				}
				if res > target {
					break
				}
			}
			if res == target {
				// fmt.Printf("%d %d\n", res, target)
				sum += target
				break
			}
		}
	}
	return sum
}

// Brute force all possible combinations of operations using trinary
// runs in about 2.9 seconds
func part2(input []string) int {
	equations := parseInput(input)
	sum := 0
	for _, eq := range equations {
		target := eq[0]
		eq = eq[1:]
		for n := range int(math.Pow(3, float64(len(eq)-1))) {
			res := eq[0]
			for i := 0; i < len(eq)-1; i++ {
				b := eq[i+1]
				op := byte((n / int(math.Pow(3, float64(i)))) % 3)
				switch op {
				case 0x0:
					res += b
				case 0x1:
					res *= b
				case 0x2:
					digits := int(math.Log10(float64(b))) + 1
					concat := int(math.Pow(10, float64(digits)))*res + b
					res = concat
				}
				if res > target { // this saved about 0.4 seconds
					break
				}
			}
			if res == target {
				sum += target
				break
			}
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

	lines := strings.Split(string(input), "\n")
	utils.Timer(func() {
		fmt.Printf("Part 1: %d\n", part1(lines))
	}, "Part 1")
	utils.Timer(func() {
		fmt.Printf("Part 2: %d\n", part2(lines))
	}, "Part 2")
}
