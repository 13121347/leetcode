package InsertSort

//https://leetcode.cn/problems/insertion-sort-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next

	for curr != nil {
		//顺序，不用插入节点，继续往后遍历
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			//逆序，需要交换cur和lastOrder节点
			prev := dummyHead
			//prev指向nonoeder 要插入地方的前一个节点，所以不能用prev.Next != lastSorted//错：找到curr节点前的节点，保存，用于链表操作（单链表只能通过遍历整个链表找前驱）
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}

		curr = lastSorted.Next
	}
	return dummyHead.Next
}
