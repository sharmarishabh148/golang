/*
The pooling pattern allows me to manage resource usage across a well defined
number of Goroutines. As explained previously, in Go pooling is not needed for
efficiency in CPU processing like at the operating system. It’s more important for
efficiency in resource usage
*/

package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Worker is recieve only
// In worker we keep on recieving data until channel is closed
func workerPools(index int, ch <-chan string, wg *sync.WaitGroup) {
	// range helps otherwise we have to use ok syntax
	//for d := range ch {
	defer wg.Done()
	/*for {
	d, ok := <-ch
	if !ok {*/
	for d := range ch {
		fmt.Printf("workerPool %d : recv'd signal : %s\n", index, d)
	}
	fmt.Printf("workerPool %d : recv'd shutdown signal\n", index)

	//}

}

// TodoTasks dummy tasks
// Channel used is sendonly.
func todoTasks(index int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	message := "data" + strconv.Itoa(index)
	ch <- message
	fmt.Println("todoTasks : ", index, " sent signal")
}

// boundedWorkPooling: In this pattern, a pool of child goroutines is created
// to service a fixed amount of work. The parent goroutine iterates over all
// work, signalling that into the pool. Once all the work has been signaled,
// then the channel is closed, the channel is flushed, and the child
// goroutines terminate.

func main() {
	t := time.Now()
	var wgworker sync.WaitGroup
	var wgtask sync.WaitGroup
	g := runtime.GOMAXPROCS(0)
	//fmt.Println("g : ", g)
	const todo = 100
	wgworker.Add(g)
	wgtask.Add(todo)

	ch := make(chan string, todo)

	for c := 0; c < g; c++ {
		go workerPools(c, ch, &wgworker)
	}

	for t := 0; t < todo; t++ {
		go todoTasks(t, ch, &wgtask)
	}

	wgtask.Wait()
	close(ch)

	// Waiting For all Workers
	wgworker.Wait()
	//time.Sleep(5 * time.Second)
	fmt.Println("Total Time Taken :", time.Since(t))

}
