package MInMoves

import (
	"sort"
)

// https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/
// 给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最小操作数(加1减1操作)
// medium

//假设数组元素都变成 xx 时，所需的移动数最少，那么 xx 需要满足什么性质呢: x = num[n/2]
/**
（1，10），如果数组中只有1和10，让这两个数变成相等的数，变成几呢？——2～9都可以，
 2：最小的移动次数是（2-1） + （10-2）=9
 3：最小的移动次数是（3-1） + （10-3）=9
 6：最小的移动次数是（6-1） + （10-6）=9
以此类推，区间是（A,B）的两个数，变成相等的X，X需要在（A，B）区间内

扩展到N个数的数组（a,b,c,d,e,f,g）, 把这些数两两组和，x尽量在所有的区间内就可以
解法：从小到大排序，x = num[n/2]

*/

func minMoves(nums []int) (ans int) {
	// 排序
	sort.Ints(nums)
	// 找到中间值
	x := nums[len(nums)/2]
	// 遍历计算其他数字变为x需要的步骤，累加
	for _, num := range nums {
		ans = ans + abs(num-x)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//2 快速选择，直接找到数组中第K个最大元素
