package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	//compare two slices with DeepEqual.
	//Check if twoSum returns {0,2} index's when given {1,2,5,7} as a sorted array
	//and 6 as a target.
	if !reflect.DeepEqual(twoSum([]int{1, 2, 5, 7}, 6), []int{0, 2}) {
		t.Error("Expected {1, 2, 5, 7} with target 6 to produce {0,2}")
	}
}

func TestTableTwoSum(t *testing.T) {
	var tests = []struct {
		sortedA  []int
		target   int
		expected []int
	}{
		{[]int{1, 4, 5, 100}, 105, []int{2, 3}},
		{[]int{100, 200, 550, 1000}, 750, []int{1, 2}},
	}

	for _, test := range tests {
		if output := twoSum(test.sortedA, test.target); !reflect.DeepEqual(output, test.expected) {
			t.Error("Failed: Expected", test.expected, "with sortedArray:", test.sortedA, "and target:", test.target)
		}
	}
}
