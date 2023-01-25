package main

import (
	"fmt"
)

func main() {
	fmt.Println("creating a channel")

	var globalString string
	// With buffered channels, the send happens before the corresponding receive.
	// The write to a happens before the send on c, which happens before the
	// corresponding receive on c completes, which happens before the print.
	c := make(chan int, 1)

	go func() {
		fmt.Println("sending on channel")
		globalString = "Hello World"
		c <- 0 // Sending Signal
		fmt.Println("sent on channel")
	}()

	fmt.Println("recieving on channel")
	<-c
	fmt.Println("recieved  on channel")

	fmt.Printf("\n\n")
	fmt.Println("globalString is : ", globalString)
}
