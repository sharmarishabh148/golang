package main

import "fmt"

//A closure is a function value that references variables from outside its body.
//The function may access and assign to the referenced variables; in this sense
//the function is "bound" to the variables.

// used to write middleware

func intSeq() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println("creating a new closure")
	nextInt2 := intSeq()
	fmt.Println(nextInt2())
}
