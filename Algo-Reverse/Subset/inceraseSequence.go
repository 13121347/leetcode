package Subset

//递增子序列 https://leetcode.cn/problems/increasing-subsequences/
// 不能破坏原数组顺序，这样的话就不能通过原地排序来解决重复元素问题——需要使用多一个数据结构来记录元素的使用情况
func increaseSequences(nums []int) (ans [][]int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	//每次递归得到的子集
	path := []int{}
	//去重辅助map
	//markedMap:=make(map[int]bool) 这里不需要全局的，和全排列的不一样
	//递归函数
	var backTrack func(index int)
	backTrack = func(index int) {
		markedMap := make(map[int]bool) //去重辅助map放在这里
		//ans要收集每一次path的值，所以不是放在终止条件里再append了,但是放进来的一个条件是，至少有两个元素
		if len(path) > 1 {
			ans = append(ans, append([]int(nil), path...))
		}

		//终止条件是index到了最后一位
		if index == len(nums) {
			return
		}

		//遍历
		for i := index; i < len(nums); i++ {
			//这里增加递增的判断
			if len(path) > 0 && path[len(path)-1] > nums[i] {
				continue
			}
			//这里增加去重的判断
			if markedMap[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			markedMap[nums[i]] = true
			backTrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	//调用函数
	backTrack(0)
	return
}
