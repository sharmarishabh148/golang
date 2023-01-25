package main

import (
	"fmt"
)

var greetings = []string{"Hi", "Hello", "Namste", "ciao"}

func main() {
	// Unbuffered chan capacity 0
	//ch := make(chan string)
	ch := make(chan string, 1)
	go greetFromSendOnlyChannel(ch)

	// In case of Buffered channel sender is able to complete
	// But reciver can wait
	//time.Sleep(5 * time.Second)

	// greeting := <-ch
	fmt.Println("Main ready")
	/*for {
		greeting, ok := doRecieve(ch)
		if !ok {
			break
		}
		fmt.Println("Greeting recieved")
		fmt.Println(greeting)
	}*/

	// can be written as
	for greeting := range ch {
		fmt.Println("Greeting recieved")
		fmt.Println(greeting)
	}

	//time.Sleep(2 * time.Second)
	fmt.Println("Main done")

}

// recieved only func greet(ch <-chan string) {
// send only func greet(ch chan<- string) {
func greet(ch chan string) {
	fmt.Printf("Greeter ready!\nGreeter waiting to send greeting...\n")
	ch <- "Hello, World!"
	fmt.Println("Greeting completed!")
}

func greetFromSendOnlyChannel(ch chan<- string) {
	fmt.Printf("Greeter ready!\n")
	for _, g := range greetings {
		ch <- g
	}
	fmt.Println("Greeting completed!")
	defer close(ch)
}
func doRecieve(ch chan string) (data string, ok bool) {
	data, ok = <-ch
	if !ok {
		fmt.Println("channel is  closed!")
		return
	}
	fmt.Println("channel is  open : ", data)
	return
}
