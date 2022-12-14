package LeafSimilar

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	vals := []int{}

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			vals = append(vals, node.Val)
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root1)
	vals1 := append([]int(nil), vals...)
	dfs(root2)
	if len(vals) != len(vals1) {
		return false
	}
	for i, v := range vals1 {
		if v != vals[i] {
			return false
		}
	}
	return true
}
