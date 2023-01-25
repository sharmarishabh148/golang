package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type MyHashSet struct {
	Size    int
	Entries []*Node
}

func Constructor() MyHashSet {
	return MyHashSet{Size: 999, Entries: make([]*Node, 999)}
}

func (this *MyHashSet) Add(key int) {
	k := key % this.Size
	n := this.Entries[k]

	if n == nil {
		this.Entries[k] = &Node{Val: key}
		return
	}

	for node := n; node != nil; node = node.Next {
		if node.Val == key {
			return
		}

		if node.Next == nil {
			node.Next = &Node{Val: key}
			return
		}
	}
}

func (this *MyHashSet) Remove(key int) {
	k := key % this.Size
	n := this.Entries[k]

	if n == nil {
		return
	}

	if n.Val == key {
		this.Entries[k] = n.Next
		return
	}

	for node := n; node != nil; node = node.Next {
		if node.Next != nil && node.Next.Val == key {
			node.Next = node.Next.Next
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	k := key % this.Size
	n := this.Entries[k]

	for node := n; node != nil; node = node.Next {
		if node.Val == key {
			return true
		}
	}

	return false
}

func main() {
	obj2 := Constructor()
	obj2.Add(1)
	obj2.Add(2)

	param := obj2.Contains(1)
	fmt.Printf("param : %v\n", param)
	param = obj2.Contains(3)
	fmt.Printf("param : %v\n", param)

	obj2.Add(2)
	param = obj2.Contains(2)
	fmt.Printf("param : %v\n", param)

	obj2.Remove(2)
	param = obj2.Contains(2)
	fmt.Printf("param : %v\n", param)

}
