package IntersectionLIst

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	countA, countB := 0, 0
	countAPtr, countBPtr := headA, headB

	for countAPtr != nil {
		countAPtr = countAPtr.Next
		countA++
	}
	for countBPtr != nil {
		countBPtr = countBPtr.Next
		countB++
	}

	movedA := headA
	movedB := headB
	if countA > countB {
		diffStep := countA - countB
		for i := 0; i < diffStep; i++ {
			movedA = movedA.Next
		}
	} else {
		diffStep := countB - countA
		for i := 0; i < diffStep; i++ {
			movedB = movedB.Next
		}
	}

	for movedA != nil && movedB != nil {
		if movedA == movedB { //出错在这里用movedA.next和movedB.next来判断
			return movedA
		}
		movedA = movedA.Next
		movedB = movedB.Next
	}
	return nil
}
