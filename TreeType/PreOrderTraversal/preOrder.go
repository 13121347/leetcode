package PreOrderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序：根左右
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	var preOrderFunc func(root *TreeNode)
	preOrderFunc = func(root *TreeNode) {
		if root == nil {
			return
		}

		result = append(result, root.Val)
		preOrderFunc(root.Left)
		preOrderFunc(root.Right)
	}
	preOrderFunc(root)
	return result
}
