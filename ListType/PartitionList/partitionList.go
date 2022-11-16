package PartitionList

//分隔链表 https://leetcode-cn.com/problems/partition-list/
//*双链表问题*
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	//特殊情况
	if head == nil || x == 0 {
		return head
	}

	//分两半遍历链表，定位第一个x数据--不行，x不一定是链表里的数，那就是要分两个指针，第二个指针指向大于x的第一个数
	dummy := &ListNode{0, head}
	prenode, backNode := dummy, head
	for backNode != nil && backNode.Val <= x {
		backNode = backNode.Next
		prenode = prenode.Next
	}

	//从preNode和backNode开始向后遍历，backnode中有小于x的，则置换到preNode后面

	for backNode != nil {
		if backNode.Val <= x {
			//置换
			prenode.Next = backNode.Next
			backNode.Next = head
			dummy.Next = backNode
			backNode = prenode.Next
			head = dummy.Next //忘记把head往前移动

		} else {
			backNode = backNode.Next
			prenode = prenode.Next
		}

	}

	return dummy.Next

}

//维护两个链表，一个存储大节点，一个存储小节点，最后把两个链表合并起来
func partitionTwoList(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small

	large := &ListNode{}
	largeHead := large

	//遍历原链表，放入大小两个链表中
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}

	//合并大小两个链表
	large.Next = nil
	small.Next = largeHead.Next

	return smallHead.Next
}
