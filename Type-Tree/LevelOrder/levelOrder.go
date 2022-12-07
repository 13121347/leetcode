package LevelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//queue
func levelOrder(root *TreeNode) [][]int {
	resultArray := make([][]int, 0)

	if root == nil {
		return resultArray
	}

	treeQueue := make([]*TreeNode, 0)
	treeQueue = append(treeQueue, root)

	for i := 0; len(treeQueue) > 0; i++ {
		levelList := make([]*TreeNode, 0)
		for j := 0; j < len(treeQueue); j++ {
			node := treeQueue[j]
			resultArray[i] = append(resultArray[i], node.Val)

			if node.Left != nil {
				levelList = append(levelList, node.Left)
			}
			if node.Right != nil {
				levelList = append(levelList, node.Right)
			}
		}
		treeQueue = levelList
	}
	return resultArray
}
