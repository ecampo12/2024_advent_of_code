package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{
			"AAAA",
			"BBCD",
			"BBCC",
			"EEEC",
		}, 140},
		{[]string{
			"RRRRIICCFF",
			"RRRRIICCCF",
			"VVRRRCCFFF",
			"VVRCCCJFFF",
			"VVVVCJJCFE",
			"VVIVCCJJEE",
			"VVIIICJJEE",
			"MIIIIIJJEE",
			"MIIISIJEEE",
			"MMMISSJEEE",
		}, 1930},
	}

	if len(tests) == 0 {
		t.Log("Not implemented")
		t.Fail() // Fail instead of skip so that the test is run every time
		return
	}

	for _, test := range tests {
		actual := part1(test.input)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		// {[]string{
		// 	"AAAA",
		// 	"BBCD",
		// 	"BBCC",
		// 	"EEEC",
		// }, 80},
		{[]string{
			"RRRRIICCFF",
			"RRRRIICCCF",
			"VVRRRCCFFF",
			"VVRCCCJFFF",
			"VVVVCJJCFE",
			"VVIVCCJJEE",
			"VVIIICJJEE",
			"MIIIIIJJEE",
			"MIIISIJEEE",
			"MMMISSJEEE",
		}, 1206},
	}

	if len(tests) == 0 {
		t.Log("Not implemented")
		t.Fail() // Fail instead of skip so that the test is run every time
		return
	}

	for _, test := range tests {
		actual := part2(test.input)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
