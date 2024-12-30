package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"utils"
)

type Gate struct{ op, left, right, output string }

var OPERATIONS = map[string]func(int, int) int{
	"OR":  func(a, b int) int { return a | b },
	"AND": func(a, b int) int { return a & b },
	"XOR": func(a, b int) int { return a ^ b },
}

func parseInput(input string) (map[string]int, []Gate) {
	wires := map[string]int{}
	gates := []Gate{}

	parts := strings.Split(input, "\n\n")
	for _, line := range strings.Split(parts[0], "\n") {
		p := strings.Split(line, ": ")
		wires[p[0]] = utils.Int(p[1])
	}

	for _, line := range strings.Split(parts[1], "\n") {
		p := strings.Split(line, " ")
		gates = append(gates, Gate{p[1], p[0], p[2], p[4]})
	}

	return wires, gates
}

func eval(wires map[string]int, gates []Gate) map[string]int {
	for {
		changed := false
		for _, gate := range gates {
			if _, ok := wires[gate.output]; ok {
				continue
			}

			left, left_ok := wires[gate.left]
			right, right_ok := wires[gate.right]

			if left_ok && right_ok {
				wires[gate.output] = OPERATIONS[gate.op](left, right)
				changed = true
			}
		}

		if !changed {
			break
		}
	}

	return wires
}

func part1(input string) int {
	wires, gates := parseInput(input)
	wires = eval(wires, gates)

	zWires := []int{}
	for i := 0; i < len(wires); i++ {
		if z, ok := wires[fmt.Sprintf("z%02d", i)]; ok {
			zWires = append(zWires, z)
		} else {
			break
		}
	}
	// reverse the wires
	for i, j := 0, len(zWires)-1; i < j; i, j = i+1, j-1 {
		zWires[i], zWires[j] = zWires[j], zWires[i]
	}
	// convert to decimal
	str := utils.Apply(zWires, func(x int) string { return strconv.Itoa(x) })
	decimal, _ := strconv.ParseInt(strings.Join(str, ""), 2, 64)
	return int(decimal)
}

func checkRippleCarry(gates []Gate) (int, utils.Set[string]) {
	ops := make(map[string]string)
	for _, gate := range gates {
		ops[gate.left+"|"+gate.right+"|"+gate.op] = gate.output
		ops[gate.right+"|"+gate.left+"|"+gate.op] = gate.output
	}
	carries := make([]string, 45)
	correct := utils.NewSet[string]()
	prevIntermediates := utils.NewSet[string]()

	for i := 0; i < 45; i++ {
		pos := fmt.Sprintf("%02d", i)
		predigit := ops["x"+pos+"|y"+pos+"|XOR"]
		precarry1 := ops["x"+pos+"|y"+pos+"|AND"]
		if i == 0 {
			carries[i] = precarry1
			continue
		}
		digit := ops[carries[i-1]+"|"+predigit+"|XOR"]
		if digit != fmt.Sprintf("z%s", pos) {
			return i - 1, correct
		}
		correct.Add(carries[i-1])
		correct.Add(predigit)
		for wire := range prevIntermediates {
			correct.Add(wire)
		}
		precarry2 := ops[carries[i-1]+"|"+predigit+"|AND"]
		carryOut := ops[precarry1+"|"+precarry2+"|OR"]
		carries[i] = carryOut
		prevIntermediates.Add(precarry1)
		prevIntermediates.Add(precarry2)
	}
	return 45, correct
}

func pairs(n int) [][2]int {
	result := make([][2]int, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			result = append(result, [2]int{i, j})
		}
	}
	return result
}

func part2(input string) string {
	_, gates := parseInput(input)
	swaps := []string{}

	base, baseUsed := checkRippleCarry(gates)
	for i := 0; i < 4; i++ {
		for _, p := range pairs(len(gates)) {
			x, y := p[0], p[1]

			if gates[x].output == "z00" || gates[y].output == "z00" {
				continue
			}

			if _, ok := baseUsed[gates[x].output]; ok {
				continue
			}

			if _, ok := baseUsed[gates[y].output]; ok {
				continue
			}

			gates[x].output, gates[y].output = gates[y].output, gates[x].output
			attempt, used := checkRippleCarry(gates)
			if attempt > base {
				swaps = append(swaps, gates[x].output, gates[y].output)
				base = attempt
				baseUsed = used
				break
			}
			gates[x].output, gates[y].output = gates[y].output, gates[x].output
		}
	}

	sort.Strings(swaps)
	return strings.Join(swaps, ",")
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
	fmt.Printf("Part 2: %s\n", part2_ans)
}
