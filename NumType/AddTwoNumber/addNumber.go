package AddTwoNumber

type ListNode struct {
	Val  int
	Next *ListNode
}

//单游标指针的移动
//进位处理

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	//异常情况
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	//正常情况,遍历l1 l2
	carry := 0
	var movedP *ListNode
	for l1 != nil || l2 != nil {
		valfroml1, valfroml2 := 0, 0
		if l1 != nil {
			valfroml1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			valfroml2 = l2.Val
			l2 = l2.Next
		}
		sum := valfroml2 + valfroml1 + carry
		//处理进位
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{sum, nil}
			movedP = head
		} else {
			movedP.Next = &ListNode{sum, nil}
			movedP = movedP.Next
		}
	}

	if carry > 0 {
		movedP.Next = &ListNode{carry, nil}
	}
	return
}
