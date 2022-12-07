package ReverseList

// 反转链表 https://leetcode.cn/problems/intersection-of-two-linked-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

//递归
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

//迭代
func reverseList2(head *ListNode) *ListNode {
	prev := &ListNode{}
	moved := head

	for moved != nil {
		nextNode := moved.Next
		moved.Next = prev
		prev = moved
		moved = nextNode
	}

	return moved //这里return的是moved，不是prev
}

//利用数组存储
func reverseListStack(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	listNodeStack := make([]*ListNode, 0)

	for head != nil {
		listNodeStack = append(listNodeStack, head)
		head = head.Next
	}

	listNodeNewHead := &ListNode{}
	movedPtr := listNodeNewHead

	for i := len(listNodeStack) - 1; i >= 0; i-- {
		movedPtr.Next = listNodeStack[i]
		movedPtr = movedPtr.Next
	}
	movedPtr.Next = nil

	return listNodeNewHead.Next
}
