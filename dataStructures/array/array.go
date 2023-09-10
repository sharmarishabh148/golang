package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}

	for i, val := range arr {
		fmt.Println("index: ", i, "value: ", val)
	}
	fmt.Println("cap: ", cap(arr), "len: ", len(arr))
	// first argument to append must be a slice; have arr (variable of
	// type [5]int)compilerInvalidAppend
	//arr = append(arr, 6)

	fmt.Println()
	for i := 0; i < len(arr); i++ {
		// In C style
		fmt.Println("index: ", i, "value: ", arr[i])
	}
}
