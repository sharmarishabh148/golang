package main

import (
	"fmt"
	"strings"
)

/*
Given an array, colors, which contains a combination of the following three elements:

0

	(representing red)

1

	(representing white)

2

	(representing blue)
*/

func sortColors(colors []int) []int {

	// Initialize the red, white, and blue pointers
	red := 0
	white := 0
	blue := len(colors) - 1

	for white <= blue {

		// If the white pointer is pointing to red
		if colors[white] == 0 {

			// Swap the values if the red pointer is not pointing to red
			if colors[red] != 0 {
				colors[red], colors[white] = colors[white], colors[red]
			}

			// Increment both the red and white pointers
			white++
			red++
		} else if colors[white] == 1 {

			// If the white pointer is pointing to white, no swapping is required
			// Just increment the white pointer
			white++
		} else { // If the white pointer is pointing to blue
			if colors[blue] != 2 {
				// Swap the values if the blue pointer is not pointing to blue
				colors[white], colors[blue] = colors[blue], colors[white]
			}

			// Decrement the blue pointer
			blue--
		}
	}

	return colors
}

// Driver code
func main() {
	inputs := [][]int{
		{0, 1, 0},
		{1, 1, 0, 2},
		{2, 1, 1, 0, 0},
		{2, 2, 2, 0, 1, 0},
		{2, 1, 1, 0, 1, 0, 2},
	}

	// Iterate over the inputs and print the sorted array for each
	for i, input := range inputs {
		fmt.Printf("%d.\tcolors: %v\n", i+1, strings.Replace(fmt.Sprint(input), " ", ", ", -1))

		sortedColors := sortColors(input)

		fmt.Printf("\n\tThe sorted array is: %v\n", strings.Replace(fmt.Sprint(sortedColors), " ", ", ", -1))

		fmt.Println(strings.Repeat("-", 100))
	}
}
