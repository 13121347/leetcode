package ViewTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉树层序遍历，输出每一层最后一个元素
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			if i == length-1 {
				res = append(res, queue[0].Val)
			}
			queue = queue[1:]
		}
	}

	return res
}
