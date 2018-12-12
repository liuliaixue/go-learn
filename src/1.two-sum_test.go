package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	if sum(2, 2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

type sumSource struct {
	src    []int
	sum    int
	target [2]int
}

func TestTwoSum(t *testing.T) {
	c := twoSum([]int{2, 3, 4}, 5)
	if c != [2]int{0, 1} {
		t.Error("index not right")
	}

	cases := []sumSource{
		{[]int{1, 2, 3, 4}, 4, [2]int{0, 2}},
		{[]int{1, 2, 3, 4}, 5, [2]int{0, 3}},
		{[]int{1, 2, 3, 4}, 7, [2]int{2, 3}},
	}

	for _, pair := range cases {
		if twoSum(pair.src, pair.sum) != pair.target {
			t.Error("index not right")
		}
	}
}
