package main

import (
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

	// Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
	// The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

	// Note:
	// Your returned answers (both index1 and index2) are not zero-based.
	// You may assume that each input would have exactly one solution and you may not use the same element twice.
	// Example:

	// Input: numbers = [2, 7, 11, 15], target = 9
	// Output: [1, 2]
	// Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.

	ans1 := answer1([]int{2, 7, 11, 15}, 9)
	log.Printf("ans1: %+v\n", ans1)
}

func answer1(nums []int, target int) *[]int {
	i := 0
	j := len(nums) - 1
	for i < j {
		tmp := nums[i] + nums[j]
		if tmp == target {
			return &[]int{i + 1, j + 1}
		} else if tmp < target {
			i++
		} else {
			j--
		}
	}
	return nil
}
