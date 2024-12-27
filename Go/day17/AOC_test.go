package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`, "4,6,3,5,6,3,5,2,1,0"},
	}

	if len(tests) == 0 {
		t.Log("Not implemented")
		t.Fail() // Fail instead of skip so that the test is run every time
		return
	}

	for _, test := range tests {
		actual := part1(test.input)
		if actual != test.expected {
			t.Errorf("%s, expected %s", actual, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`, 117440},
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
