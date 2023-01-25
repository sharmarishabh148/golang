package main

import (
	"fmt"
)

var hellos = []string{"Hi", "Hello", "Namste", "ciao"}
var goodbyes = []string{"go", "bye", "GoodBye", "Jao"}

func main() {
	// Unbuffered chan capacity 0
	//ch := make(chan string)
	ch := make(chan string, 1)
	ch2 := make(chan string, 1)
	go greetFromSendOnlyChannel(hellos, ch)
	go greetFromSendOnlyChannel(goodbyes, ch2)

	// greeting := <-ch
	fmt.Println("Main ready")
	for {
		select {
		case msg1, ok := <-ch:
			if !ok {
				ch = nil
				break
			}
			fmt.Println("recieved from channel ", msg1)
		case msg2, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Println("recieved from channel2 ", msg2)
		default:
			return
		}
	}
	//time.Sleep(2 * time.Second)
	fmt.Println("Main done")

}

func greetFromSendOnlyChannel(greetings []string, ch chan<- string) {
	fmt.Printf("Greeter ready!\n")
	for _, g := range greetings {
		ch <- g
	}
	close(ch)
	fmt.Println("Greeting completed!")
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

func doWork(ch1, ch2 chan string) {
	for {
		select {
		case msg1, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Println("recieved from channel ", msg1)
		case msg2, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Println("recieved from channe2 ", msg2)
		default:
			fmt.Println("Nothing")
		}
	}
}
