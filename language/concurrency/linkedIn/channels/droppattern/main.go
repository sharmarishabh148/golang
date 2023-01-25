package main

import (
	"fmt"
	"time"
)

func drop() {
	const cap = 100
	ch := make(chan string, cap)
	go func() {
		for p := range ch {
			fmt.Println("child : recv'd signal :", p)
		}
	}()
	const work = 2000
	dropped := 0
	for w := 0; w < work; w++ {
		// A select blocks until one of its cases can run, then it executes that case.
		// It chooses one at random if multiple are ready.
		select {
		case ch <- "data":
			fmt.Println("parent : sent signal :", w)
		default:
			dropped++
			fmt.Println("parent : dropped data :", w)
			time.Sleep(10 * time.Microsecond)
		}
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")
	time.Sleep(time.Second)
	fmt.Printf("-----------------dropped = %d--------------------", dropped)
	fmt.Println("-------------------------------------------------")
}

func main() {
	drop()
}
