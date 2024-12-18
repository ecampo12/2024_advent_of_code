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
			"p=0,4 v=3,-3",
			"p=6,3 v=-1,-3",
			"p=10,3 v=-1,2",
			"p=2,0 v=2,-1",
			"p=0,0 v=1,3",
			"p=3,0 v=-2,-2",
			"p=7,6 v=-1,-3",
			"p=3,0 v=-1,-2",
			"p=9,3 v=2,3",
			"p=7,3 v=-1,2",
			"p=2,4 v=2,-3",
			"p=9,5 v=-3,-3",
		}, 12},
	}

	if len(tests) == 0 {
		t.Log("Not implemented")
		t.Fail() // Fail instead of skip so that the test is run every time
		return
	}

	for _, test := range tests {
		actual := part1(test.input, true)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
