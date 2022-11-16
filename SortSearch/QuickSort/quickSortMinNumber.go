package QuickSort

import (
	"strconv"
	"strings"
)

// 把数组排成最小的数 https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
// 非负整数数组拼接
// 自定义大小+快排
func minNumber(nums []int) (result string) {
	quickSort(nums, 0, len(nums)-1)
	for _, v := range nums {
		vStr := strconv.Itoa(v)
		result = result + vStr
	}
	return result
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	partitionIndex := partition(nums, left, right)
	quickSortMinNumber(nums, left, partitionIndex-1)
	quickSortMinNumber(nums, partitionIndex+1, right)
}

func quickSortMinNumber(nums []int, left, right int) int {
	key, index := nums[left], nums[left+1]

	for j := left + 1; j < right; j++ {
		if lessThan(nums[j], key) {
			nums[j], nums[index] = nums[index], nums[j]
		}
	}

	nums[index-1], nums[left] = nums[left], nums[index-1]
	return index - 1
}

//自定义大小比较函数
func lessThan(x, y int) bool {
	xStr := strconv.Itoa(x)
	yStr := strconv.Itoa(y)
	return strings.Compare(xStr, yStr) < 0
}
