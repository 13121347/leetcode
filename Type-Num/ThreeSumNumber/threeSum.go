package ThreeSumNumber

import (
	"sort"
)

//三指针 https://leetcode.cn/problems/3sum/
//easy 首先对数组进行排序，然后遍历数组
//然后再在当前节点后面取左右指针，判断左右指针的值是否等于0-nums[i]，然后分别左右移动
//怎么去重呢? 满足条件时，看左指针的值是否和前一个位置相等，右指针的值是否和和它后一个位置的值相等
//排序+双指针
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	ans := make([][]int, 0)
	//遍历最外层
	for first := 0; first < n; first++ {
		//去重
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// third =外层右边结束的元素
		third := n - 1
		target := -1 * nums[first]
		// 遍历最外层右边开始的元素 b
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue //去重
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				continue
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
