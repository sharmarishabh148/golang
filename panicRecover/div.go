package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println((safeDiv(1, 0)))
	fmt.Println((safe(7, 2)))
}

func safeDiv(a, b int) (q int, err error) {
	// Using Named Return Values
	defer func() {
		if e := recover(); e != nil {
			log.Println("Error", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	q = a / b
	return q, nil
}

func safe(a, b int) int {
	return a / b
}
