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
			"............",
			"........0...",
			".....0......",
			".......0....",
			"....0.......",
			"......A.....",
			"............",
			"............",
			"........A...",
			".........A..",
			"............",
			"............",
		}, 14},
	}

	if len(tests) == 0 {
		t.Log("Not implemented")
		t.Fail() // Fail instead of skip so that the test is run every time
		return
	}

	for _, test := range tests {
		actual := part1(test.input)
		if actual != test.expected {
			t.Errorf("%d expected %d", actual, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{
			"............",
			"........0...",
			".....0......",
			".......0....",
			"....0.......",
			"......A.....",
			"............",
			"............",
			"........A...",
			".........A..",
			"............",
			"............",
		}, 34},
	}

	if len(tests) == 0 {
		t.Log("Not implemented")
		t.Fail() // Fail instead of skip so that the test is run every time
		return
	}

	for _, test := range tests {
		actual := part2(test.input)
		if actual != test.expected {
			t.Errorf("%d expected %d", actual, test.expected)
		}
	}
}
