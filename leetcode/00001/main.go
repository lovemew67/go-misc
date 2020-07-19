package main

import (
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	// https://leetcode.com/problems/two-sum/submissions/

	// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	// Note:
	// You may assume that each input would have exactly one solution, and you may not use the same element twice.
	// Example:

	// Input: nums = [2, 7, 11, 15], target = 9
	// Output: [0, 1]
	// Explanation: nums[0] + nums[1] = 2 + 7 = 9.

	ans1 := answer1([]int{2, 7, 11, 15}, 9)
	ans2 := answer2([]int{2, 7, 11, 15}, 9)
	ans3 := answer3([]int{2, 7, 11, 15}, 9)
	log.Printf("ans1: %+v\n", ans1)
	log.Printf("ans2: %+v\n", ans2)
	log.Printf("ans3: %+v\n", ans3)
}

// time:  n^2
// space: 1
func answer1(nums []int, target int) *[]int {
	log.Printf("nums: %+v\n", nums)
	log.Printf("target: %d\n", target)
	for i := 0; i < (len(nums) / 2); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return &[]int{i, j}
			}
		}
	}
	return nil
}

// time:  n
// space: n
func answer2(nums []int, target int) *[]int {
	numMap := map[int]int{}
	for index := 0; index < len(nums); index++ {
		numMap[nums[index]] = index
	}
	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]
		if v, ok := numMap[sub]; ok {
			return &[]int{i, v}
		}
	}
	return nil
}

// time:  n
// space: n
func answer3(nums []int, target int) *[]int {
	numMap := map[int]int{}
	for i, num := range nums {
		numMap[num] = i
		if index, ok := numMap[target-num]; ok {
			return &[]int{index, i}
		}
	}
	return nil
}
