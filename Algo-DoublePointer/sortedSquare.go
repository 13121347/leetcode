package Algo_DoublePointer

// 有序数组的平方 https://leetcode.cn/problems/squares-of-a-sorted-array/
// easy 这个数组在取平方之前，是有序的，那么它绝对值最大的数一定是在两端的
//所以我们可以定义两个指针，一个指向最左端，一个指向最右端，比较两者平方的大小，大的平方放入结果数组，并移动指针

func sortedSquares(nums []int) []int {
	left, right, newIndex := 0, len(nums)-1, 0
	newArr := make([]int, len(nums))
	leftValue := nums[left] * nums[left]
	rightValue := nums[right] * nums[right]

	for left <= right {
		if leftValue < rightValue {
			newArr[newIndex] = rightValue
			right--
			rightValue = nums[right] * nums[right]
		} else {
			newArr[newIndex] = leftValue
			left++
			leftValue = nums[left] * nums[left]
		}
		newIndex++
	}
	return newArr
}
