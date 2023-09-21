package main

/*
Given the two distinct integer arrays, nums1 and nums2, where nums1 is a subset of nums2, find all the next greater
elements for nums1 values in the corresponding places of nums2.
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	stack := make([]int, 0, len(nums2))
	m := make(map[int]int)

	// iterate over nums2
	for _, current := range nums2 {
		// while stack is not empty and current element is greater than the top element of the stack
		for len(stack) > 0 && current > stack[len(stack)-1] {
			// update the map with the current element as the value for the popped element
			m[stack[len(stack)-1]] = current
			// pop the top element from the stack
			stack = stack[:len(stack)-1]
		}
		// push the current element to the stack
		stack = append(stack, current)
	}

	// iterate over remaining elements in the stack, pop them and set their values to -1 in the map
	for _, element := range stack {
		m[element] = -1
	}

	ans := make([]int, len(nums1))
	// iterate over nums1 and add the corresponding value from the map to ans
	for i, num := range nums1 {
		ans[i] = m[num]
	}

	return ans
}

/*
	arrayToString is used to convert an integer array to string.

arrayToString is used as an helper function in a main function for
printing purposes.
*/
func arrayToString(nums []int) string {
	res := "["
	for i, num := range nums {
		res += strconv.Itoa(num)
		if i < len(nums)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}

// Driver code
func main() {
	nums1 := [][]int{{2, 4}, {3, 2, 5}, {14, 45, 52}, {1, 3, 2}, {4, 2}, {0}}
	nums2 := [][]int{{1, 2, 3, 4}, {2, 3, 5, 1}, {52, 14, 45, 65}, {1, 3, 2, 4, 5}, {1, 2, 4, 3}, {0}}

	for i := range nums1 {
		fmt.Printf("%d.\tNums 1 = %s\n", i+1, arrayToString(nums1[i]))
		fmt.Printf("\tNums 2 = %s\n", arrayToString(nums2[i]))
		result := nextGreaterElement(nums1[i], nums2[i])
		fmt.Printf("\n\tThe Next Greater Element Array = %s\n", arrayToString(result))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
