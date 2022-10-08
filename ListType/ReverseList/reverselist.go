package ReverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

func reverseList2(head *ListNode) *ListNode {
	prev := &ListNode{}
	curr := head

	for curr != nil {
		tempNode := curr.Next

		curr.Next = prev
		prev = curr
		curr = tempNode
	}

	return prev
}
