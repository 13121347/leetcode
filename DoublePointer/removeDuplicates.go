package DoublePointer

// 原地删除数组重复数
func removeDuplicateElement(nums []int) int {
	after, front := 1, 1
	for ; front < len(nums); front++ {
		if nums[front] != nums[front-1] {
			nums[after] = nums[front]
			after++
		}
	}
	return after
}
