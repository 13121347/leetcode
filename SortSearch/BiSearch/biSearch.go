package BiSearch

// 二分查找 https://leetcode.cn/problems/binary-search/
// easy
// 时间复杂度：O(log⁡n) 空间复杂度：O(1)O(1)O(1)

func biSearch(nums []int, target int) int {
	leftIdx, rightIdx := 0, len(nums)-1

	for leftIdx <= rightIdx {
		midIndex := leftIdx + (rightIdx-leftIdx)/2 //结果上看起来等于（left+right/2,但是这样可能会溢出
		if nums[midIndex] == target {
			return midIndex
		} else if nums[midIndex] > target {
			rightIdx = midIndex - 1
		} else {
			leftIdx = midIndex + 1
		}
	}
	return -1
}

func biSearchRound2(nums []int, target int) int {
	leftIndex, rightIndex := 0, len(nums)-1

	for leftIndex <= rightIndex {
		midIndex := leftIndex + (rightIndex-leftIndex)/2
		if target == nums[midIndex] {
			return midIndex
		} else if target > nums[midIndex] {
			leftIndex = midIndex + 1
		} else {
			rightIndex = midIndex - 1
		}
	}
	return -1
}
