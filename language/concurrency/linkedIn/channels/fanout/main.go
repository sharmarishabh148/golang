package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fanout() {
	tasks := 2000

	ch := make(chan string, tasks)

	for t := 0; t < tasks; t++ {
		go func(emp int) {
			// Simulation of unknown task
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "signal"
			fmt.Printf("employee task : %d sent signal\n", emp)
		}(t)

	}

	for tasks > 0 {
		p := <-ch
		tasks--
		//fmt.Println(p)
		fmt.Printf("manager task : %d recv'd %s\n", tasks, p)
	}
}

func main() {
	fanout()
}
