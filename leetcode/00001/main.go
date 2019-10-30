package main

import (
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 00001_two-sum")
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
}
