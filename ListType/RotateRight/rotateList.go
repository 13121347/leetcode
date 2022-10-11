package RotateRight

type ListNode struct {
	Val  int
	Next *ListNode
}

//用map，新的下标计算复杂
func rotateRight(head *ListNode, k int) *ListNode {
	//异常情况
	if head == nil || k == 0 {
		return head
	}

	nodesMapping := make(map[int]*ListNode)
	headPtr := head
	idx := 1
	for headPtr != nil {
		nodesMapping[idx] = headPtr
		headPtr = headPtr.Next
		idx++
	}

	newNodesMapping := make(map[int]*ListNode)
	length := idx - 1
	//重新计算下标
	steps := k % length
	for originIdx, node := range nodesMapping {
		if originIdx+steps > length {
			newIdx := steps - (length - originIdx)
			newNodesMapping[newIdx] = node
		} else {
			newIdx := originIdx + steps
			newNodesMapping[newIdx] = node
		}

	}

	//恢复成list
	dummy := &ListNode{}
	movedPtr := dummy

	for i := 1; i <= idx; i++ {
		movedPtr.Next = newNodesMapping[i]
		movedPtr = movedPtr.Next
	}

	return dummy.Next
}

// 链表
func rotateRightList(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	nodes := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	length := len(nodes)

	steps := k % length
	if steps == 0 {
		return head
	}

	// 不重新生成一个list，直接在原list上操作-----相对顺序是不变的
	index := length - steps
	lastNode := nodes[index-1]
	lastNode.Next = nil

	movedNode := nodes[index]
	cur = movedNode
	for i := 1; i < steps; i++ {
		movedNode = movedNode.Next
	}
	movedNode.Next = head
	return cur
}

//双指针
func rotateRightListDouble(head *ListNode, k int) *ListNode {
	//nil情况
	if head == nil {
		return head
	}

	//长度计算
	countPtr := head
	lentgh := 0
	for countPtr != nil {
		countPtr = countPtr.Next
		lentgh++
	}

	//快慢指针找出分割点
	steps := k % lentgh
	if steps == 0 {
		return head
	}

	fast, slow := head, head
	//快指针先走K步
	for i := 0; i < steps; i++ {
		fast = fast.Next
	}
	//快慢指针同时往下走，直到快指针走到最后,此时慢指针指向新list的最后一个node
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//切割原list
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head

	return newHead
}
