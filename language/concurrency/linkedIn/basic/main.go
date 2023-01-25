package main

import "fmt"

func main() {
	go hello()
	goodbye()
}

func hello() {
	fmt.Println("Hello, World")
}
func goodbye() {
	fmt.Println("GoodBye, World")
}
