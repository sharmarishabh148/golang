package main

import (
	"fmt"
	"strings"
)

func removeNthLastNode(head *EduLinkedListNode, n int) *EduLinkedListNode {

	right := head
	left := head

	for i := 0; i < n; i++ {
		right = right.next
	}

	if right == nil {
		return head.next
	}

	for right.next != nil {
		right = right.next
		left = left.next
	}

	left.next = left.next.next

	return head
}

// Driver Code
func main() {
	inputs := [][]int{
		{23, 89, 10, 5, 67, 39, 70, 28},
		{34, 53, 6, 95, 38, 28, 17, 63, 16, 76},
		{288, 224, 275, 390, 4, 383, 330, 60, 193},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{69, 8, 49, 106, 116, 112, 104, 129, 39, 14, 27, 12},
	}

	n := []int{4, 1, 6, 9, 11}

	for i := 0; i < len(inputs); i++ {
		inputLinkedList := EduLinkedList{}
		inputLinkedList.CreateLinkedList(inputs[i])
		fmt.Printf("%d.\tLinked List:\t\t", i+1)
		DisplayLinkedListWithForwardArrow(inputLinkedList.head)
		fmt.Printf("\n\tn = %d", n[i])
		fmt.Printf("\n\tUpdated Linked List:\t")
		DisplayLinkedListWithForwardArrow(removeNthLastNode(inputLinkedList.head, n[i]))
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}