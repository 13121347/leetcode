package ReverseBetween

// 反转一段链表 https://leetcode-cn.com/problems/reverse-linked-list-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	//异常情况
	if head == nil || left == right {
		return head
	}

	//所有节点存到数组中
	listNodeStack := make([]*ListNode, 0)
	headPtr := head
	for headPtr != nil {
		listNodeStack = append(listNodeStack, headPtr)
		headPtr = headPtr.Next
	}

	resultList := &ListNode{}
	movedPtr := resultList

	//左半不动
	for i := 0; i < left-1; i++ {
		movedPtr.Next = listNodeStack[i]
		movedPtr = movedPtr.Next
	}

	//反转中间部分
	for i := right - 1; i >= left-1; i-- {
		movedPtr.Next = listNodeStack[i]
		movedPtr = movedPtr.Next
	}

	//右半不动
	for i := right; i < len(listNodeStack); i++ {
		movedPtr.Next = listNodeStack[i]
		movedPtr = movedPtr.Next
	}
	movedPtr.Next = nil

	return resultList.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head

	for cur != nil {
		nextNode := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextNode
	}
	return cur
}

//时间复杂度：O(N)
func reverseListBetween(head *ListNode, left, right int) *ListNode {
	// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// 第 1 步：从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
	// 建议写在 for 循环里，语义清晰
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 第 2 步：从 pre 再走 right - left + 1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 第 3 步：切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	// 注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 第 4 步：同第 206 题，反转链表的子区间
	reverseList(leftNode)

	// 第 5 步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next

}

// 穿针法
func reverseBetween2(head *ListNode, left, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
