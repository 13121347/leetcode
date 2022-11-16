package BiSearch

// 搜索插入位置 https://leetcode.cn/problems/search-insert-position/
// easy

func searchInsert(nums []int, target int) int {
	leftIdx, rightIdx := 0, len(nums)-1

	for leftIdx <= rightIdx {
		mid := leftIdx + (rightIdx-leftIdx)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			rightIdx = mid - 1
		} else {
			leftIdx = mid + 1
		}
	}
	return leftIdx
}
