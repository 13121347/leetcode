package Subset

import "sort"

// 返回子集 https://leetcode.cn/problems/subsets-ii/
// 包含重复元素，结果集不可重复------事前排序，遍历时跳过nums[i]==nunms[i-1]的情况
func subsetWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	if nums == nil || len(nums) == 0 {
		return
	}

	//每次递归得到的子集
	path := []int{}

	//递归函数
	var backTrack func(index int)
	backTrack = func(index int) {
		//ans要收集每一次path的值，所以不是放在终止条件里再append了
		ans = append(ans, append([]int(nil), path...))
		//终止条件是index到了最后一位
		if index == len(nums) {
			return
		}

		//遍历
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backTrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	//调用函数
	backTrack(0)
	return
}
