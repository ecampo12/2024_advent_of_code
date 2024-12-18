package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Helper function to calculate the modulus of a negative number
func mod(a, b int) int {
	return (((a % b) + b) % b)
}

type Point struct {
	x, y int
}

type Robot struct {
	x, y   int
	vx, vy int
}

func (r *Robot) move(times, width, height int) {
	r.x = mod(r.x+times*r.vx, width)
	r.y = mod(r.y+times*r.vy, height)
}

func parseInput(input []string) []Robot {
	Robots := []Robot{}
	for _, line := range input {
		matches := regexp.MustCompile(`(-?\d+)`).FindAllString(line, -1)
		nums := make([]int, len(matches))
		for i, match := range matches {
			fmt.Sscanf(match, "%d", &nums[i])
		}
		x, y, vx, vy := nums[0], nums[1], nums[2], nums[3]
		Robots = append(Robots, Robot{x, y, vx, vy})
	}
	return Robots
}

func part1(input []string, test bool) int {
	Width, Height := 101, 103
	if test {
		Width, Height = 11, 7
	}
	Robots := parseInput(input)
	qudrants := make([]int, 4)

	halfWidth, halfHeight := Width/2, Height/2

	for _, r := range Robots {
		r.move(100, Width, Height)
		if r.x == halfWidth || r.y == halfHeight {
			continue
		}
		if r.x < halfWidth {
			if r.y < halfHeight {
				qudrants[0]++
			} else {
				qudrants[1]++
			}
		} else {
			if r.y < halfHeight {
				qudrants[2]++
			} else {
				qudrants[3]++
			}
		}
	}
	return qudrants[0] * qudrants[1] * qudrants[2] * qudrants[3]
}

// There is a way to solve this by looking at the safety factor of the robots,
// but that approach does not seem to work for my input. So I just brute force.
func part2(input []string) int {
	Width, Height := 101, 103
	Robots := parseInput(input)
	t := 0
	for true {
		t++
		pos := make(map[Point]bool)
		valid := true
		for _, r := range Robots {
			x := mod(r.x+t*r.vx, Width)
			y := mod(r.y+t*r.vy, Height)
			if pos[Point{x, y}] {
				valid = false
				break
			}
			pos[Point{x, y}] = true
		}
		if valid {
			break
		}
	}
	return t
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	part1_ans := part1(lines, false)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
