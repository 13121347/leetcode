package ReverseBetween

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
