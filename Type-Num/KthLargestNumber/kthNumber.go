package KthLargestNumber

import (
	"math/rand"
	"sort"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(arr []int, leftIndex, rightIndex, targetIndex int) int {
	q := partition2(arr, leftIndex, rightIndex)
	if q == targetIndex {
		return arr[q]
	} else if q < targetIndex {
		return quickSelect(arr, q+1, rightIndex, targetIndex)
	}
	return quickSelect(arr, leftIndex, q-1, targetIndex)
}

func randomPartition(a []int, leftIndex, rightIndex int) int {
	splitIndex := rand.Int()%(rightIndex-leftIndex+1) + 1
	a[splitIndex], a[rightIndex] = a[rightIndex], a[splitIndex]
	return partition(a, leftIndex, rightIndex)
}

func partition(a []int, leftIndex, rightIndex int) int {
	key, idx := a[rightIndex], leftIndex-1
	for j := leftIndex; j < rightIndex; j++ {
		if a[j] <= key {
			idx++
			a[idx], a[j] = a[j], a[idx]
		}
	}
	a[idx+1], a[rightIndex] = a[rightIndex], a[idx+1]
	return idx + 1
}

func partition2(a []int, leftIndex, rightIndex int) int {
	key, idx := a[leftIndex], leftIndex+1
	for j := leftIndex + 1; j < rightIndex; j++ {
		if a[j] <= key {
			a[idx], a[j] = a[j], a[idx]
			idx++
		}
	}
	a[idx-1], a[leftIndex] = a[leftIndex], a[idx-1]
	return idx - 1
}

//https://blog.csdn.net/csdn_kou/article/details/104166148?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7ERate-1-104166148-blog-118792892.pc_relevant_3mothn_strategy_recovery&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromBaidu%7ERate-1-104166148-blog-118792892.pc_relevant_3mothn_strategy_recovery&utm_relevant_index=2

func findKthLargestSort(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
