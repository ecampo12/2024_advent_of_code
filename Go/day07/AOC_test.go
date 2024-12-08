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
			"190: 10 19",
			"3267: 81 40 27",
			"83: 17 5",
			"156: 15 6",
			"7290: 6 8 6 15",
			"161011: 16 10 13",
			"192: 17 8 14",
			"21037: 9 7 18 13",
			"292: 11 6 16 20"}, 3749},
	}

	if len(tests) == 0 {
		t.Log("Not implemented")
		t.Fail() // Fail instead of skip so that the test is run every time
		return
	}

	for _, test := range tests {
		actual := part1(test.input)
		if actual != test.expected {
			t.Errorf("%d, expected %d", actual, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{
			"190: 10 19",
			"3267: 81 40 27",
			"83: 17 5",
			"156: 15 6",
			"7290: 6 8 6 15",
			"161011: 16 10 13",
			"192: 17 8 14",
			"21037: 9 7 18 13",
			"292: 11 6 16 20"}, 11387},
	}

	if len(tests) == 0 {
		t.Log("Not implemented")
		t.Fail() // Fail instead of skip so that the test is run every time
		return
	}

	for _, test := range tests {
		actual := part2(test.input)
		if actual != test.expected {
			t.Errorf("%d, expected %d", actual, test.expected)
		}
	}
}
