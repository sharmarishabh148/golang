package main

import "fmt"

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func main() {
	sq, cu := powerSeries(4)
	fmt.Println("Sqaure: ", sq, "Cube: ", cu)
}
