package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	const goRoutinesCount = 2

	var wg sync.WaitGroup
	wg.Add(goRoutinesCount)

	for g := 0; g < goRoutinesCount; g++ {
		go func() {
			for i := 0; i < 2; i++ {
				value := counter
				value++
				// add logs for ctx switch
				fmt.Println("logging")
				counter = value
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)
}
