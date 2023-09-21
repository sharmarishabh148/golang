package main

/*
For any positive number in base 10, return the complement of its binary
representation as an integer in base 10.
*/
import (
	"fmt"
	"math"
	"strings"
)

// pow calculates the power of the given digit
func pow(digit int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}

func findBitwiseComplement(num int) int {
	if num == 0 {
		return 1
	}
	bitCount := math.Floor(math.Log2(float64(num))) + 1
	allBitsSet := pow(2, int(bitCount)) - 1

	return num ^ allBitsSet
}

func main() {
	decimalValues := []int{42, 233, 100, 999999, 54}

	for i, value := range decimalValues {
		fmt.Printf("%d.\tInput: %d\n", i+1, value)
		fmt.Printf("\tBitwise complement of %d is: %d\n", value, findBitwiseComplement(value))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
