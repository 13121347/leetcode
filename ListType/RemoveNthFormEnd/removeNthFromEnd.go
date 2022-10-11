package RemoveNthFormEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	if n == 0 {
		return head
	} //少了这一种情况

	headCount := 0
	headCountPtr := head
	for headCountPtr != nil {
		headCount++
		headCountPtr = headCountPtr.Next
	}

	if n == headCount {
		return head.Next
	} //少了这一种情况，如果用了哑节点，就不需要额外处理这个逻辑

	movedPtr := head
	prePtr := &ListNode{0, head}

	locToDelete := headCount - n
	for i := 0; i < locToDelete; i++ {
		movedPtr = movedPtr.Next
		prePtr = prePtr.Next
	}

	prePtr.Next = movedPtr.Next

	return head
}

// 只用一个指针就可以了，可以进行数据删除，不用俩这么麻烦～
func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

// 栈的思想
func removeNthFromEndStack(head *ListNode, n int) *ListNode {
	modeStack := make([]*ListNode, 0)
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		modeStack = append(modeStack, node)
	}

	prev := modeStack[len(modeStack)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

//双指针 在不预处理出链表的长度，以及使用常数空间的前提下解决本题
/**
由于我们需要找到倒数第 n 个节点，因此我们可以使用两个指针first 和 second 同时对链表进行遍历，并且first 比second 超前 n 个节点。
当first遍历到链表的末尾时， second 就恰好处于倒数第  n 个节点。
*/

func removeNthFromEndTwoPtr(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy

	for i := 0; i < n; i++ {
		first = first.Next
	}

	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next

	return dummy.Next
}
