package main

import (
	"fmt"
	"slices"
)

// input: a sorted array
// goal: find the two values that equal the target
// output: return the index's of the two values that equal the target
func twoSum(nums []int, target int) []int {
	slices.Sort(nums)
	l := 0
	r := len(nums) - 1

	for l < r {
		sum := nums[l] + nums[r]
		if sum > target {
			r -= 1
		} else if sum < target {
			l += 1
		} else {
			returnValue := []int{l, r}
			return returnValue
		}

	}

	return nums
}

func main() {
	input := []int{1, 2, 5, 7}
	fmt.Println(twoSum(input, 6))
}
