package LowestConnonAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	if find(root.Left, p) && find(root.Left, q) {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if find(root.Right, p) && find(root.Right, q) {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func find(root, node *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Val == node.Val {
		return true
	}
	return find(root.Left, node) || find(root.Right, node)
}
