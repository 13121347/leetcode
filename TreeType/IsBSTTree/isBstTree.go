package IsBSTTree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST2(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if lower >= root.Val || upper <= root.Val {
		return false
	}

	return helper(root.Left, lower, root.Val) || helper(root.Right, root.Val, upper)
}
