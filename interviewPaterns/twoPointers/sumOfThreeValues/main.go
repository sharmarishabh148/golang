package main

/*
Given an array of integers, nums, and an integer value, target, determine if
there are any three integers in nums whose sum is equal to the target, that
is, nums[i] + nums[j] + nums[k] == target. Return TRUE if three such integers
exist in the array. Otherwise, return FALSE.
*/

import (
	"fmt"
	"sort"
	"strings"
)

// FindSumOfThree is our challenge function
func findSumOfThree(nums []int, target int) bool {
	// Sort the input vector
	sort.Ints(nums)

	// Create two pointers to track our indices and, a variable to store our triple sum
	low, high, triple := 0, 0, 0

	// Fix one integer at a time and find the other two
	for i := 0; i < len(nums)-2; i++ {

		// Initialize the two pointers
		low = i + 1
		high = len(nums) - 1

		// Traverse the vector to find the triplet whose sum equals the target
		for low < high {
			triple = nums[i] + nums[low] + nums[high]

			// The sum of the triplet equals the target
			if triple == target {
				return true
			} else if triple < target {
				// The sum of the triplet is less than target, so move the low pointer forward
				low += 1
			} else {
				// The sum of the triplet is greater than target, so move the high pointer backward
				high -= 1
			}
		}
	}
	// No such triplet found whose sum equals the given target
	return false
}

// Driver code
func main() {
	numsLists := [][]int{
		{3, 7, 1, 2, 8, 4, 5},
		{-1, 2, 1, -4, 5, -3},
		{2, 3, 4, 1, 7, 9},
		{1, -1, 0},
		{2, 4, 2, 7, 6, 3, 1},
	}
	tList := []int{10, 7, 20, -1, 8}
	for i, nList := range numsLists {
		fmt.Printf("%d. Input array: %s\n", i+1, strings.Replace(fmt.Sprint(nList), " ", ", ", -1))

		if findSumOfThree(nList, tList[i]) {
			fmt.Printf("   Sum for %d exists\n", tList[i])
		} else {
			fmt.Printf("   Sum for %d does not exist\n", tList[i])
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
