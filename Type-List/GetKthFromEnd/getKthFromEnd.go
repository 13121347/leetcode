package GetKthFromEnd

//获得倒数第K个节点https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	//特殊情况
	if head == nil || k == 0 {
		return nil
	}
	//正常逻辑
	fast, slow := head, head

	for i := 1; i <= k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
