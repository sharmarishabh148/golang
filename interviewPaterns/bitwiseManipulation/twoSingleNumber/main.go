package main

import (
	"fmt"
	"strings"
)

func twoSingleNumbers(arr []int) []int {
	bitwiseSum := 0
	for _, i := range arr {
		bitwiseSum ^= i
	}

	// The least significant bit can be found with number ^ (-number)
	bitwiseMask := bitwiseSum & (-bitwiseSum)

	// Divide into two groups of numbers, here we want the group with bit set
	// which results in one of the numbers we want due to the property X ^ X = 0
	results := 0
	for _, i := range arr {
		if bitwiseMask&i != 0 {
			results ^= i
		}
	}
	return []int{results, bitwiseSum ^ results}
}

// Driver code
func main() {
	inputs := [][]int{
		{1, 2, 2, 3},
		{4, 4, 3, 2, 3, 5},
		{1, 1, 7, 4, 5, 5, 8, 8},
		{1, 0},
		{9, 8, 8, 7, 6, 6, 4, 4},
	}

	for i, input := range inputs {
		fmt.Printf("%d.\tInput list: %s\n", i+1, strings.Replace(fmt.Sprint(input), " ", ", ", -1))
		result := twoSingleNumbers(input)
		fmt.Printf("\tTwo Singles numbers in a list: %s\n", strings.Replace(fmt.Sprint(result), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
