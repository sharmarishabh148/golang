/*
channels : It’s important to think of a channel not as a data structure, but as a mechanic for
signaling. This goes in line with the idea that I send and receive from a channel, not
read and write.

Unbuffered Channels:
Guarantees at the signaling level with the receive happening before send. Sending
and receiving Goroutines need to come together in the same space and time for a
signal to be processed.

Buffered Channels:
Guarantees outside of the signaling level with the send happening before the
receive. If the buffer is not full, sends can complete else they block. If the buffer is
not empty, receives can complete else they block.

Closed Channels:
A channel can be in a closed state by using the built-in function close. I don’t need
to close a channel to release memory, this is for changing the state. Sending on a
closed channel will cause a panic, however receiving on a closed channel will return
immediately

*/

// Sample program to show the order of channel communication for unbuffered,
// buffered and closing channels based on the specification.
// https://golang.org/ref/mem#tmp_7
package main

import "fmt"

func main() {
	unBuffered()
	buffered()
	closed()
}

// With unbuffered channels, the receive happens before the corresponding send.
// The write to a happens before the receive on c, which happens before the
// corresponding send on c completes, which happens before the print.
func unBuffered() {
	c := make(chan int)
	var a string

	go func() {
		a = "hello, world unBuffered"
		<-c
	}()

	c <- 0

	// We are guaranteed to print "hello, world".
	fmt.Println(a)
}

// With buffered channels, the send happens before the corresponding receive.
// The write to a happens before the send on c, which happens before the
// corresponding receive on c completes, which happens before the print.
func buffered() {
	c := make(chan int, 10)
	var a string

	go func() {
		a = "hello, world buffered"
		c <- 0
	}()

	<-c

	// We are guaranteed to print "hello, world".
	fmt.Println(a)
}

// With both types of channels, a close happens before the corresponding receive.
// The write to a happens before the close on c, which happens before the
// corresponding receive on c completes, which happens before the print.
func closed() {
	c := make(chan int, 10)
	var a string

	go func() {
		a = "hello, world closed"
		close(c)
	}()

	<-c

	// We are guaranteed to print "hello, world".
	fmt.Println(a)
}
