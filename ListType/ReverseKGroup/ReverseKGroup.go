package ReverseKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		//tail向后走K步，到达需要反转的一段List尾部
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
			//保留下一段list头部
			nex := tail.Next
			//反转head-tail，即第N小段
			head, tail = myReverse(head, tail)
			//连接N小段和前段
			pre.Next = head
			//连接N小段和后段
			tail.Next = nex
			//准备下一段的处理
			pre = tail
			head = tail.Next
		}
		return hair.Next
	}

}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
