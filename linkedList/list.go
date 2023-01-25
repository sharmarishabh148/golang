package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func findLen(head *ListNode) int {
	l := 0
	for ; head != nil; head = head.Next {
		l++
	}
	return l
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	len1, len2 := findLen(headA), findLen(headB)
	if len1 > len2 { // list A is bigger
		// traverse difference node in List A
		for ; len1 > len2; len1-- {
			headA = headA.Next
		}
	} else {
		// traverse difference node in List A
		for ; len2 > len1; len2-- {
			headB = headB.Next
		}
	}
	// Notice we are checking for pointer addresses
	for headA != headB {
		headA, headB = headA.Next, headB.Next
	}
	return headA
}

func deleteDuplicates(head *ListNode) *ListNode {
	// since the list is alredy sorted
	for curr := head; curr != nil; {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}
	return head
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func getDecimalValue(head *ListNode) int {
	r := 0
	for head != nil {
		r = (r << 1) | head.Val
		head = head.Next
	}
	return r
}

func removeElements(head *ListNode, val int) *ListNode {
	var dummy *ListNode = &ListNode{Next: head}
	for curr, iterator := dummy, head; iterator != nil; {
		if iterator.Val == val {
			curr.Next = iterator.Next // update current's  Next
			iterator = iterator.Next
		} else {
			curr = curr.Next // move current and iterator
			iterator = iterator.Next
		}
	}
	return dummy.Next
}
func reverseList(head *ListNode) *ListNode {
	var prev, curr *ListNode
	curr = head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// even sized length linked list function
// For Odd size list return slow ptr
func endOfFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)
	fmt.Println("firstHalfEnd :", firstHalfEnd.Val)
	fmt.Println("secondHalfStart :", secondHalfStart.Val)
	frontPosition := head
	midPosition := secondHalfStart
	result := true

	for result && (midPosition != nil) {
		if frontPosition.Val != midPosition.Val {
			fmt.Println("secondHalfStart :", secondHalfStart.Val)
			firstHalfEnd.Next = reverseList(secondHalfStart)
			fmt.Println("newHead :", head.Val)
			return false
		}
		frontPosition = frontPosition.Next
		midPosition = midPosition.Next
	}
	fmt.Println("secondHalfStart :", secondHalfStart.Val)
	firstHalfEnd.Next = reverseList(secondHalfStart)
	fmt.Println("newHead :", head.Val)

	fmt.Println()
	return result

}

func (ll *LinkedList) insertAtBeginning(data interface{}) {
	newNode := &ListNode{
		// Type Assertions
		Val: data.(int),
	}
	if ll.head == nil {
		ll.head = newNode
	} else {
		newNode.Next = ll.head
		ll.head = newNode
	}
	ll.size++
}

func (ll *LinkedList) display() error {
	if ll.head == nil {
		return fmt.Errorf("display: List is empty")
	}
	current := ll.head
	for current != nil {
		if current.Next == nil {
			fmt.Printf("%v", current.Val)
		} else {
			fmt.Printf("%v -> ", current.Val)
		}

		current = current.Next
	}
	fmt.Println()
	return nil
}

func main() {
	ll := LinkedList{}
	fmt.Printf("insertAtBeginning: 5\n")
	ll.insertAtBeginning(5)
	fmt.Printf("insertAtBeginning: 4\n")
	ll.insertAtBeginning(4)
	fmt.Printf("insertAtBeginning: 3\n")
	ll.insertAtBeginning(3)
	fmt.Printf("insertAtBeginning: 2\n")
	ll.insertAtBeginning(2)
	fmt.Printf("insertAtBeginning: 1\n")
	ll.insertAtBeginning(1)
	fmt.Printf("insertAtBeginning: 0\n")
	ll.insertAtBeginning(0)

	fmt.Printf("Displaying LinkedList: \n")
	ll.display()
	res := isPalindrome(ll.head)
	fmt.Printf("isPalindrome :%v\n", res)
	ll.display()

	ll2 := LinkedList{}
	ll2.insertAtBeginning(5)
	ll2.insertAtBeginning(4)
	ll2.insertAtBeginning(3)
	ll2.insertAtBeginning(2)
	ll2.insertAtBeginning(1)
	ll2.insertAtBeginning(1)
	fmt.Printf("Linked List 2 :\n")
	ll2.display()
	ll2.head = removeElements(ll2.head, 1)
	ll2.display()

	fmt.Printf("middleNode of Linked List 2 :\n")
	mn := middleNode(ll2.head)
	fmt.Printf(":%v\n\n", mn.Val)

	ll3 := LinkedList{}
	ll3.insertAtBeginning(0)
	ll3.insertAtBeginning(1)
	ll3.insertAtBeginning(1)
	ll3.insertAtBeginning(0)
	ll3.insertAtBeginning(1)
	ll3.insertAtBeginning(0)
	ll3.insertAtBeginning(1)
	ll3.display()

	h := &ListNode{Val: 3}
	h.Next = &ListNode{Val: 2}
	h.Next.Next = &ListNode{Val: 0}
	h.Next.Next.Next = &ListNode{Val: -4,
		Next: h.Next}
	//Next: nil}
	//cycleList := &LinkedList{head: h}
	fmt.Printf("cycle present in cycleList  : %v\n", hasCycle(h))
	// cycleList.display() infite loop
	fmt.Printf("Decimal Equivalent of Linked List3  : %v\n", getDecimalValue(ll3.head))

	ll4 := LinkedList{}
	ll4.insertAtBeginning(3)
	ll4.insertAtBeginning(3)
	ll4.insertAtBeginning(2)
	ll4.insertAtBeginning(1)
	ll4.insertAtBeginning(1)
	ll4.display()
	fmt.Printf("deleteDuplicates Linked List4  : %v\n", deleteDuplicates(ll4.head))
	ll4.display()

	fmt.Printf("Modifying Linked List4  : \n")

	intersectionPoint := ll4.head
	ll4.insertAtBeginning(0)
	ll4.insertAtBeginning(-1)
	ll4.display()

	ylist := &ListNode{Val: 11}
	ylist.Next = &ListNode{Val: 12}
	ylist.Next.Next = &ListNode{Val: 13,
		Next: intersectionPoint}
	YNodeList := &LinkedList{head: ylist}
	YNodeList.display()

	if getIntersectionNode(ll4.head, YNodeList.head) == intersectionPoint {
		fmt.Printf("getIntersectionNode for YNodeList ,Linked List4  : True\n")
	}
	fmt.Printf("getIntersectionNode for YNodeList ,Linked List4  : %v\n", getIntersectionNode(ll4.head, YNodeList.head).Val)

	fmt.Printf("HashMap testing\n")
}
