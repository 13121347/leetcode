package Dynamic

import (
	"fmt"
	"sort"
)

// 最长递增子序列 https://leetcode.cn/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) (ans int) {
	//处理边界情况
	if len(nums) == 0 {
		return
	}
	//定义dp[i]，表示前i个数字里，最长的子串长度
	dp := make([]int, len(nums))

	//初始化数据，dp[0]=0,dp[1]=1
	dp[0] = 1
	maxans := 1
	//状态转移方程 dp[i]=max(dp[j])+1,maxLength = max(dp[i])
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxans = max(maxans, dp[i])
	}
	fmt.Println(dp)
	return maxans
}

//这样做不正确的原因是，状态转移方程有问题
//之前做到的题，状态转移方程都是dp[i]和dp[i-1]/dp[i-2]这样的关系，即当前这个元素选/不选依赖于相邻的那个元素
//但是现在这个递增序列是不连续的，当前的值选择/不选择依赖于上一个选择的数值，而不是相邻的那个数
//所以这里要两边遍历，像方法一那样
func lengthOfLIS2(nums []int) (ans int) {
	//处理边界情况
	if len(nums) == 0 {
		return
	}
	//定义dp[i]，表示前i个数字里，最长的子串长度，正确的定义是dp[i]: 以第i个为结尾的
	//什么时候是前i，什么时候是第i呢
	dp := make([]int, len(nums)+1)

	//初始化数据，dp[0]=0,
	dp[0] = 0
	dp[1] = 1
	//状态转移方程 dp[i]=max(dp[j])+1,maxLength = max(dp[i])
	for i := 2; i <= len(nums); i++ {
		if nums[i-1] < nums[i-2] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	fmt.Print(dp)
	sort.Ints(dp)
	return dp[len(dp)-1]
}
