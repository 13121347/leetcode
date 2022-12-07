package Algo_DoublePointer

// 移除元素  https://leetcode.cn/problems/remove-element/
// easy 右指针指向当前将要处理的元素，左指针指向下一个要赋值的位置
//时间复杂度：O(n)
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { //双指针，不一定就是首尾俩指针
		if v != val {
			nums[left] = v
			left++
		}
	}

	return left
}

//优化
func removeElementOptimize(nums []int, val int) int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}

	return left
}
