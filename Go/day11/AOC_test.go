package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		blinks   int
	}{
		{"125 17", 22, 6},
		{"125 17", 55312, 25},
	}

	if len(tests) == 0 {
		t.Log("Not implemented")
		t.Fail() // Fail instead of skip so that the test is run every time
		return
	}

	for _, test := range tests {
		actual := part1(test.input, test.blinks)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
