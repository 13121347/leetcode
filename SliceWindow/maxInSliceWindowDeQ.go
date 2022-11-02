package SliceWindow

func maxSlidingWindowDQ(nums []int, k int) []int {
	//q 记录下标
	q := []int{}

	// 如果nums[i] > nums[q最后一个]，则弹出q,把i放进去，这样，q中的index对应的num就是单调递增的
	// 这里的for循环，比较新的数据和Q的最后一个数据大小，新数据大于Q最后一个数，则弹出。保证q[0]是这个队列里最大的数
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	//初始化q
	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]

	for i := k; i < n; i++ {
		push(i)
		//q[0]不在当前窗口内，即在窗口左边了
		for q[0] <= i-k {
			//相当于双端队列，弹出q[0]
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
