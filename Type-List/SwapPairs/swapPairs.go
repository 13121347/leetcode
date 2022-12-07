package SwapPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// nil和长度为1时，直接返回
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	dummyHead := &ListNode{0, head}
	moved := dummyHead

	for moved.Next != nil && moved.Next.Next != nil {
		node1 := moved.Next
		node2 := moved.Next.Next

		moved.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		moved = node1
	}

	return dummyHead.Next
}
