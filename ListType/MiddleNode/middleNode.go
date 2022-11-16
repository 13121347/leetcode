package MiddleNode

// 链表中间节点https://leetcode.cn/problems/middle-of-the-linked-list/
// 简单

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil { //这里要注意判断条件
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
