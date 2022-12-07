package BiSearch

//  在排序数组中查找元素的第一个和最后一个位置 https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	//左边届
	leftBound := leftBound(nums, target)
	rightBound := rightBound(nums, target)

	if rightBound < leftBound {
		return []int{-1, -1}
	}
	return []int{leftBound, rightBound}
}

func leftBound(nums []int, target int) int {
	leftIdx, rightIdx := 0, len(nums)-1

	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)>>1
		//左移
		if target <= nums[midIdx] { //这里是小于等于
			rightIdx = midIdx - 1
		} else {
			//右移
			leftIdx = midIdx + 1
		}
	}
	return leftIdx
}

func rightBound(nums []int, target int) int {
	leftIdx, rightIdx := 0, len(nums)-1

	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)>>1
		if target >= nums[midIdx] {
			leftIdx = midIdx + 1
		} else {
			rightIdx = midIdx - 1
		}
	}
	return rightIdx
}
