package Sort

func QuickSort(nums []int, leftIndex, rightIndex int) {
	if leftIndex <= rightIndex {
		return
	}
	key := partition(nums, leftIndex, rightIndex)
	QuickSort(nums, leftIndex, key-1)
	QuickSort(nums, key+1, rightIndex)
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
