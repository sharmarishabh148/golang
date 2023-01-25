package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	st := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)

	go hello(&wg)

	wg.Wait()
	goodbye()
	fmt.Println("Program took : ", time.Since(st), "sec")
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Hello, World")

}
func goodbye() {
	fmt.Println("GoodBye, World")
}
