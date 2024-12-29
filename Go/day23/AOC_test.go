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
			"kh-tc",
			"qp-kh",
			"de-cg",
			"ka-co",
			"yn-aq",
			"qp-ub",
			"cg-tb",
			"vc-aq",
			"tb-ka",
			"wh-tc",
			"yn-cg",
			"kh-ub",
			"ta-co",
			"de-co",
			"tc-td",
			"tb-wq",
			"wh-td",
			"ta-ka",
			"td-qp",
			"aq-cg",
			"wq-ub",
			"ub-vc",
			"de-ta",
			"wq-aq",
			"wq-vc",
			"wh-yn",
			"ka-de",
			"kh-ta",
			"co-tc",
			"wh-qp",
			"tb-vc",
			"td-yn"}, 7},
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
		expected string
	}{
		{[]string{
			"kh-tc",
			"qp-kh",
			"de-cg",
			"ka-co",
			"yn-aq",
			"qp-ub",
			"cg-tb",
			"vc-aq",
			"tb-ka",
			"wh-tc",
			"yn-cg",
			"kh-ub",
			"ta-co",
			"de-co",
			"tc-td",
			"tb-wq",
			"wh-td",
			"ta-ka",
			"td-qp",
			"aq-cg",
			"wq-ub",
			"ub-vc",
			"de-ta",
			"wq-aq",
			"wq-vc",
			"wh-yn",
			"ka-de",
			"kh-ta",
			"co-tc",
			"wh-qp",
			"tb-vc",
			"td-yn"}, "co,de,ka,ta"},
	}

	if len(tests) == 0 {
		t.Log("Not implemented")
		t.Fail() // Fail instead of skip so that the test is run every time
		return
	}

	for _, test := range tests {
		actual := part2(test.input)
		if actual != test.expected {
			t.Errorf("%s, expected %s", actual, test.expected)
		}
	}
}
