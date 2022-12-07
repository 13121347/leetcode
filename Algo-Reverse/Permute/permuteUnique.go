package Permute

import "sort"

//唯一的全排列 https://leetcode.cn/problems/permutations-ii/
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums) //1 多了排序
	length := len(nums)
	marked := make([]bool, length)
	perm := []int{}
	var backTrack func(int)
	backTrack = func(idx int) {
		if idx == length {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if marked[i] || i > 0 && !marked[i-1] && v == nums[i-1] { //2 多了判断减枝条件
				continue
			}
			perm = append(perm, v)
			marked[i] = true
			backTrack(idx + 1)
			marked[i] = false
			perm = perm[:len(perm)-1]
		}
	}

	backTrack(0)
	return
}
