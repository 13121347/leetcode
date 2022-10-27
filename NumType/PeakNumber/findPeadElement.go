package PeakNumber

import (
	"math"
	"math/rand"
)

func findPeakElement(nums []int) (idx int) {
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	return
}

func findPeakElement2(nums []int) (idx int) {
	n := len(nums)
	idx = rand.Intn(n)

	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	for !(get(idx-1) < get(idx) && get(idx) > get(idx+1)) {
		if get(idx) < get(idx+1) {
			idx++
		} else {
			idx--
		}
	}
	return idx
}

func findPeakElement3(nums []int) (idx int) {
	n := len(nums)

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, n-1
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
