package Combination

// 输出逆序对 https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
// 超时了

func reversePairs(nums []int) (res int) {
	//记录picked 2 num
	path := make([]int, 0)

	var backTrack func(index int)
	backTrack = func(index int) {

		//终止条件
		if len(path) == 2 { //这里的终止条件不是index >= len(nums) || len(path)>2{ ,不用对index做这个限制
			if path[0] > path[1] {
				res++
			}
			return
		}
		//遍历
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			backTrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backTrack(0)
	return
}
