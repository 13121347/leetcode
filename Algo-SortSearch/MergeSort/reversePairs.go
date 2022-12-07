package MergeSort

var count int

// 输出逆序对的个数 https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func reversePairs(nums []int) (res int) {
	res = countWhenMergeSort(nums, 0, len(nums)-1)
	return
}

func countWhenMergeSort(nums []int, start, end int) (count int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	countWhenMergeSort(nums, start, mid)
	countWhenMergeSort(nums, mid+1, end)
	countWhenMerge(nums, start, mid, end, count)
	return
}

func countWhenMerge(nums []int, start, mid, end, count int) {
	leftPartIndex, rightPartIndex, tmpIndex := start, mid+1, 0
	tmpArr := make([]int, end-start+1)

	//开始merge
	for leftPartIndex <= mid && rightPartIndex <= end {
		if nums[leftPartIndex] <= nums[rightPartIndex] {
			tmpArr[tmpIndex] = nums[leftPartIndex]
			leftPartIndex++
			tmpIndex++
		} else {
			tmpArr[tmpIndex] = nums[rightPartIndex]
			rightPartIndex++
			tmpIndex++
			count += (mid - leftPartIndex + 1)
		}
	}

	//处理剩余数据
	if leftPartIndex <= mid {
		tmpArr[tmpIndex] = nums[leftPartIndex]
		tmpIndex++
		leftPartIndex++
	}

	if rightPartIndex <= end {
		tmpArr[tmpIndex] = nums[rightPartIndex]
		tmpIndex++
		rightPartIndex++
	}

	//将临时数组的元素拷贝给原数组
	for k, v := range tmpArr {
		nums[start+k] = v
	}
}
