// Send and receive integers to a channel using go routines.

// In the main function call a function "generator" which will add 10
//integers to a channel; return the channel from the functions.
// In the main function call a function "receiver"; pass the channel to this
// function and print all the 10 integers.

package main

import (
	"fmt"
	"sync"
)

var sum int

func generator(data int, c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// send data from here to reciver
	// send to channel
	c <- data
	fmt.Println("=================================================>")
	fmt.Printf("data :%d send from generator\n", data)
	fmt.Println("------------------------------------------------->")

}

// recives channel
func receiver(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// reciev from chanel
	data := <-c
	sum += data
	fmt.Println("<=================================================")
	fmt.Printf("data :%d received at receiver\n", data)
	fmt.Println("<-------------------------------------------------")
}

func question1() {
	// Unbuffered channel ,
	c := make(chan int)
	// for concurrency
	var wgGenerator sync.WaitGroup
	var wgReceiver sync.WaitGroup
	count := 10
	wgGenerator.Add(count)
	wgReceiver.Add(count)
	for index := 0; index < count; index++ {
		go generator(index, c, &wgGenerator)
	}

	for count > 0 {
		count--
		go receiver(c, &wgReceiver)

	}

	wgGenerator.Wait()
	close(c)
	wgReceiver.Wait()
	//time.Sleep(time.Second)
	fmt.Println("----------------------sum is---------------------", sum)
}

func generating10Numbers(num []int) <-chan int {
	c := make(chan int, 10)
	for _, n := range num {
		c <- n
	}
	close(c)
	return c
}

func receiving10Numbers(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range c {
		fmt.Println(n)
	}
	fmt.Println("Sending stop signal no more data to read")
}
func question2() {
	var wg sync.WaitGroup
	ch := generating10Numbers([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	go receiving10Numbers(ch, &wg)
	wg.Wait()
	fmt.Println("Done waiting")
}
func main() {
	//question1()
	question2()
}
