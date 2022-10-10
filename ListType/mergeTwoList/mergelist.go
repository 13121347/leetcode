package mergeTwoList

type ListNode struct {
	Val  int
	Next *ListNode
}

//暴力，这里重新创建了一个新的list，把l1，l2的值复制过来，内存使用增加了
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 处理特殊情况：list出现nil
	if list1 == nil && list2 == nil {
		return nil
	}
	var head, movedPtr *ListNode
	list1moved := list1
	list2moved := list2

	currentVal := 0
	// 遍历两个list，按升序拼接
	for list1moved != nil && list2moved != nil {
		if list1moved.Val >= list2moved.Val {
			currentVal = list2moved.Val
			list2moved = list2moved.Next
		} else {
			currentVal = list1moved.Val
			list1moved = list1moved.Next
		}
		if head == nil {
			head = &ListNode{currentVal, &ListNode{}}
			movedPtr = head
		} else {
			movedPtr.Next = &ListNode{currentVal, &ListNode{}}
			//最开始复习链表，总是不记得要把moved指针往后移动
			movedPtr = movedPtr.Next
		}
	}

	// 其中一个list已经处理到最后
	if list1moved == nil {
		for list2moved != nil {
			if head == nil {
				head = &ListNode{list2moved.Val, &ListNode{}}
				movedPtr = head
			} else {
				movedPtr.Next = &ListNode{list2moved.Val, &ListNode{}}
				movedPtr = movedPtr.Next
			}
			list2moved = list2moved.Next
		}
	}
	if list2moved == nil {
		for list1moved != nil {
			if head == nil {
				head = &ListNode{list1moved.Val, &ListNode{}}
				movedPtr = head
			} else {
				movedPtr.Next = &ListNode{list1moved.Val, &ListNode{}}
				movedPtr = movedPtr.Next
			}
			list1moved = list1moved.Next
		}
	}

	movedPtr.Next = nil

	return head
}

// 暴力优化，优化了重复定义的内容
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	// 处理特殊情况：list出现nil
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil && list2 != nil {
		return list2
	}

	if list2 == nil && list1 != nil {
		return list1
	}

	var head, movedPtr *ListNode
	list1moved := list1
	list2moved := list2

	currentVal := 0
	// 遍历两个list，按升序拼接
	for list1moved != nil && list2moved != nil {
		if list1moved.Val >= list2moved.Val {
			currentVal = list2moved.Val
			list2moved = list2moved.Next
		} else {
			currentVal = list1moved.Val
			list1moved = list1moved.Next
		}
		if head == nil {
			head = &ListNode{currentVal, &ListNode{}}
			movedPtr = head
		} else {
			movedPtr.Next = &ListNode{currentVal, &ListNode{}}
			movedPtr = movedPtr.Next
		}
	}

	// 其中一个list已经处理到最后
	if list1moved != nil {
		movedPtr.Next = list1moved
		return head
	}
	if list2moved != nil {
		movedPtr.Next = list2moved
		return head
	}

	movedPtr.Next = nil
	return head
}

// 哑节点暴力，对于这种在原链表上操作的，一般都是定义一个哑节点
func mergeTwoListsDummy(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyListNode := &ListNode{}
	movedP := dummyListNode

	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			movedP.Next = l2
			l2 = l2.Next
		} else {
			movedP.Next = l1
			l1 = l1.Next
		}
		movedP = movedP.Next
	}

	if l1 != nil {
		movedP.Next = l1
		return dummyListNode.Next
	}
	if l2 != nil {
		movedP.Next = l2
		return dummyListNode.Next
	}
	movedP.Next = nil
	return dummyListNode.Next
}

//针对上面，如果不用哑节点，这样做就比较麻烦了，因为不确定第一个节点应该指向l1/l2,所以要比较多的特殊处理
func mergeTwoListsDummy2(l1 *ListNode, l2 *ListNode) *ListNode {
	movedL1 := l1
	movedL2 := l2

	var movedHead *ListNode
	//定头节点在哪里
	if movedL1.Val <= movedL2.Val {
		movedHead = movedL1
		//遍历L1剩下部分，和L2

	} else {
		movedHead = movedL2
		//遍历L2剩下部分，和L1
	}

	return movedHead

}

// 递归
func mergeTwoLists3(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists3(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists3(list1, list2.Next)
		return list2
	}
}
