package SortSearch

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
