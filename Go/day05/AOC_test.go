package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`, 143},
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
		input    string
		expected int
	}{
		{`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`, 123},
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
