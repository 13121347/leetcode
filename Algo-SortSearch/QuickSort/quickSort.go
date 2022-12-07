package QuickSort

/**
快排，每一次排序固定一个数值的位置
平摊期望时间是 O(nlogn)
快速排序不能保证相等元素的相对顺序不发生改变，所以不稳定
https://blog.csdn.net/Hell_potato777/article/details/113698205
*/

//升序
func QuickSortUp(nums []int, leftIndex, rightIndex int) {
	if leftIndex > rightIndex {
		return
	}
	keyIdx := partition(nums, leftIndex, rightIndex)
	QuickSortUp(nums, leftIndex, keyIdx-1)
	QuickSortUp(nums, keyIdx+1, rightIndex)
}

//一趟排序，获取分界值的下标
func partition(a []int, leftIndex, rightIndex int) int {
	//分界值为第一个元素，idx为第二个元素
	key, idx := a[leftIndex], leftIndex+1
	for j := leftIndex + 1; j <= rightIndex; j++ {
		if a[j] < key {
			a[idx], a[j] = a[j], a[idx]
			idx++
		}
	}
	a[idx-1], a[leftIndex] = a[leftIndex], a[idx-1]
	return idx - 1
}

//降序
func QuickSortDown(nums []int, leftIndex, rightIndex int) {
	if leftIndex > rightIndex {
		return
	}
	key := partionDown(nums, leftIndex, rightIndex)
	QuickSortDown(nums, leftIndex, key-1)
	QuickSortDown(nums, key+1, rightIndex)
}

func partionDown(a []int, leftIndex, rightIndex int) int {
	key, idx := a[leftIndex], leftIndex+1

	for j := leftIndex + 1; j <= rightIndex; j++ {
		if a[j] >= key {
			a[j], a[idx] = a[idx], a[j]
			idx++
		}
	}
	a[leftIndex], a[idx-1] = a[idx-1], a[leftIndex]
	return idx - 1
}

//或者使用内置排序
