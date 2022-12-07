package RemoveNode

//https://leetcode.cn/problems/remove-linked-list-elements/
//虚拟头节点删除法，将头节点和其他节点的处理归依化 prev节点并不是每次都移动的
//时间复杂度O（n）,空间复杂度O（1）

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	//1 处理特殊情况
	if head == nil {
		return head
	}

	//2 遍历list进行删除
	//2.1 定义pre节点用于返回
	dummy := &ListNode{Next: head}
	//2.2 定义move节点用于遍历
	cur := head
	prev := dummy
	//2.3从头开始遍历
	for cur != nil {
		//执行删除动作
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}

	return dummy.Next
}
