package midoflinkedlist

type EduLinkedList struct {
	head *EduLinkedListNode
}

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

/*
InsertNodeAtHead method will insert a LinkedListNode at head

	of a linked list.
*/
func (l *EduLinkedList) InsertNodeAtHead(node *EduLinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

// Function to find the middle node of the linked list
func GetMiddleNode(head *EduLinkedListNode) *EduLinkedListNode {

	// Create two pointers, slow and fast ,initially pointing to the head
	slow := head
	fast := head

	// Traverse the linked list until fast reaches at the last node or NULL
	for fast != nil && fast.next != nil {

		// Move the slow pointer one step ahead
		slow = slow.next

		// Move the fast pointer two steps ahead
		fast = fast.next.next
	}

	// Return the slow pointer
	return slow
}

func NewLinkedListNode(data int, next *EduLinkedListNode) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = nil
	return node
}
