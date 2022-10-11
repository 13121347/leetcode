package lastOrderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//左 右 根
func lastOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	var lastOrder func(root *TreeNode)
	lastOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		lastOrder(root.Left)
		lastOrder(root.Right)
		result = append(result, root.Val)
	}
	lastOrder(root)

	return result

}
