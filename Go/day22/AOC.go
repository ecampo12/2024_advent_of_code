package main

import (
	"fmt"
	"os"
	"strings"
	"utils"
)

func nextNum(x int) int {
	x ^= (x * 64) % 16777216
	x ^= (x / 32) % 16777216
	x ^= (x * 2048) % 16777216
	return x
}

func part1(input []string) int {
	nums := utils.Apply(input, utils.Int)

	return utils.ApplySum(nums, func(x int) int {
		for i := 0; i < 2000; i++ {
			x = nextNum(x)
		}
		return x
	})
}

type Sequence struct{ a, b, c, d int }

func getSequences(price []int, change []int) map[Sequence]int {
	seq := map[Sequence]int{}
	for i := 0; i < len(change)-3; i++ {
		pattern := Sequence{change[i], change[i+1], change[i+2], change[i+3]}
		if _, ok := seq[pattern]; !ok {
			seq[pattern] = price[i+4]
		}
	}
	return seq
}

func part2(input []string) int {
	nums := utils.Apply(input, utils.Int)
	prices := map[int][]int{}
	scores := map[Sequence]int{}
	for _, x := range nums {
		key := x
		prices[key] = make([]int, 2000)
		for i := 0; i < 2000; i++ {
			x = nextNum(x)
			prices[key][i] = x % 10
		}
		price := prices[key]
		change := []int{}
		for i := 0; i < len(price)-1; i++ {
			change = append(change, price[i+1]-price[i])
		}
		seq := getSequences(price, change)
		for k, v := range seq {
			scores[k] += v
		}
	}
	max := 0
	for _, v := range scores {
		if v > max {
			max = v
		}
	}
	return max
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
