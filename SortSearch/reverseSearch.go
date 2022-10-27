package SortSearch

func reverseSearch(nums []int, target int) int {
	leftIdx, rightIdx := 0, len(nums)-1

	for leftIdx <= rightIdx {
		midIndex := leftIdx + (rightIdx-leftIdx)/2
		if target == nums[midIndex] {
			return midIndex
		}
		if nums[midIndex] < nums[rightIdx] {
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
