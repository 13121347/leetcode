package SliceWindow

// 爱生气的书店老板 https://leetcode.cn/problems/grumpy-bookstore-owner/
// medium 不生气的部分作为固定窗口，固定窗口把customers分成了三部分，最后求三部分的最大和

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	withinSum, leftSum, rightSum := 0, 0, 0

	len := len(customers)
	//窗口位于起点
	for i := 0; i < minutes; i++ {
		withinSum = withinSum + customers[i]
	}
	//窗口位于起点时右区间的值
	for i := minutes; i < len; i++ {
		if grumpy[i] == 0 {
			rightSum = rightSum + customers[i]
		}
	}
	//窗口左右-开始移动窗口
	left, right := 1, minutes
	maxSum := withinSum + leftSum + rightSum

	//移动，重新计算左中右区间的值
	for right < len {
		//重新计算左区间的值
		if grumpy[left-1] == 0 {
			leftSum = leftSum + customers[left-1]
		}
		//重新计算右区间的值
		if grumpy[right] == 0 {
			rightSum = rightSum - customers[right]
		}
		//窗口值
		withinSum = withinSum - customers[left-1] + customers[right]
		//最大和
		maxSum = max(maxSum, leftSum+withinSum+rightSum)
		left++
		right++
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
