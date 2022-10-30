package Permute

//全排列 https://leetcode.cn/problems/permutations/

func permute(nums []int) (ans [][]int) {
	mp := make(map[int]int)
	tmp := make([]int, 0)

	var dfs func(now int)
	dfs = func(now int) {
		if now == len(nums) {
			ans = append(ans, append([]int(nil), tmp...))
			return
		}
		for _, j := range nums {
			if mp[j] != 1 {
				mp[j] = 1
				tmp = append(tmp, j)
				dfs(now + 1)
				tmp = tmp[:len(tmp)-1]
				mp[j] = 0
			}
		}
	}

	dfs(0)
	return ans
}

func permute2(nums []int) [][]int {
	ans := make([][]int, 0)
	var dfs func(now int)
	dfs = func(now int) {
		if now == len(nums) {
			ans = append(ans, append([]int(nil), nums...))
		}
		for j := now; j < len(nums); j++ {
			nums[j], nums[now] = nums[now], nums[j]
			dfs(now + 1)
			nums[j], nums[now] = nums[now], nums[j]
		}
	}
	dfs(0)
	return ans
}
