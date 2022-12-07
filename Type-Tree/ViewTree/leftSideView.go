package ViewTree

//二叉树层序遍历，输出每一层最后一个元素
func leftSideView(root *TreeNode) []int {
	//处理边界情况
	if root == nil {
		return nil
	}
	//定义返回数组
	var res []int
	//层序遍历，定义queue队列,入队root节点
	queue := []*TreeNode{root}

	//循环处理队列中的数据
	for len(queue) > 0 {
		currentLevelNodeNm := len(queue)
		//按层处理node
		for i := 0; i < currentLevelNodeNm; i++ {
			//从左到右入队下一层node
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			//记录最左边的节点
			if i == 0 {
				res = append(res, queue[0].Val)
			}
			//弹出已处理过的node
			queue = queue[1:]
		}
	}
	return res
}
