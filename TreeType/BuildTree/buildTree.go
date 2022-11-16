package BuildTree

//从前序和中序遍历序列构造二叉树
//https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

/**
前序遍历： 【根】【左子树】【右子树】
中序遍历： 【左子树】【根】【右子树】
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}

	//找到数的根节点
	i := 0
	for ; i < len(preorder); i++ {
		//前序遍历第一个节点=这棵数树的根节点，这一步就是找根节点
		if inorder[i] == preorder[0] {
			break
		}
	}

	//递归求在被根节点index划分的左右两棵树
	//中序： 0....左....i....右....len(inorder)-1
	//前序： 0.............len(inorder[:1])...........len(preoeder)-1
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}
