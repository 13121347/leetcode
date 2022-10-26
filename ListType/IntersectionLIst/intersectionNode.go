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
	//这里最开始没有定义countAPtr,countBPtr,直接用head向下遍历，导致遍历完count，head指针已经是nil了
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
		//出错在这里用movedA.next和movedB.next来判断。要明确指针指在哪里了
		if movedA == movedB {
			return movedA
		}
		movedA = movedA.Next
		movedB = movedB.Next
	}
	return nil
}

//double pointer
func getIntersectionNodeDB(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa
}
