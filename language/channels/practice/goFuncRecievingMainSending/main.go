package main

import (
	"fmt"
)

func main() {
	fmt.Println("creating a channel")
	// Unbuffered channel

	var globalString string
	c := make(chan int)

	go func() {
		fmt.Println("recieving on channel")
		<-c
		globalString = "Hello World"
		fmt.Println("recieved on channel")
	}()

	fmt.Println("sending on channel")
	c <- 0
	fmt.Println("sent on channel")

	fmt.Printf("\n\n")
	fmt.Println("globalString is : ", globalString)
}
