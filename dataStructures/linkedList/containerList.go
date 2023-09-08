package main

import (
	"container/list"
	"fmt"
)

func main() {
	// initializing zero valued list.
	var intList list.List
	l := list.New()
	e := l.PushFront("a")

	fmt.Printf("%+v :\n", e)
	fmt.Printf("%+v :\n", l)

	fmt.Println()
	// pushing some data in int
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for item := intList.Front(); item != nil; item = item.Next() {
		fmt.Print(item.Value.(int), " ")
	}

}
