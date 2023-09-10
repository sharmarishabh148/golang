package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	fmt.Println("cap: ", cap(slice), "len: ", len(slice))
	slice = append(slice, 6)

	fmt.Println("cap: ", cap(slice), "len: ", len(slice))

	fmt.Println("2D slices")

	row := 5
	col := 5
	var sliceOfSlices = make([][]int, row)

	// slices has tow range for loop
	// Using Pointer symantics i is a variable that takes on the
	// values of the indices of the elements in the someSlice.
	for i := range sliceOfSlices { // Traversing row only
		sliceOfSlices[i] = make([]int, col)
	}

	fmt.Println(sliceOfSlices)

	// Remember slices is a 3 word DataStructure
}
