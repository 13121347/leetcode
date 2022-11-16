package ReOrderList

// 重排链表 https://leetcode.cn/problems/reorder-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

//时间复杂度O(N)
//空间复杂度O（N）
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	//数组保存node
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	//重建list
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
