package BiSearch

//搜索旋转排序数组 https://leetcode.cn/problems/search-in-rotated-sorted-array/
//给定一个旋转后的数组，查某个数
func reverseSearch(nums []int, target int) int {
	leftIdx, rightIdx := 0, len(nums)-1

	for leftIdx <= rightIdx {
		midIndex := leftIdx + (rightIdx-leftIdx)/2
		//如果target数据 = nums[mid],直接返回
		if target == nums[midIndex] {
			return midIndex
		}
		//target != mid，如果nums[mid] < nums[right]，说明mid到right部分是有序的
		if nums[midIndex] < nums[rightIdx] {
			//根据target大小，判断下一次在哪个部分search
			if target > nums[midIndex] && target <= nums[rightIdx] {
				leftIdx = midIndex + 1
			} else {
				rightIdx = midIndex - 1
			}
		} else { //前半部分有序，在前半段判断
			if target >= nums[leftIdx] && target < nums[midIndex] {
				rightIdx = midIndex - 1
			} else {
				leftIdx = midIndex + 1
			}
		}
	}
	return -1
}
