package KthLargestNumber

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(arr []int, leftIndex, rightIndex, targetIndex int) int {
	q := partition(arr, leftIndex, rightIndex)
	if q == targetIndex {
		return arr[q]
	} else if q < targetIndex {
		return quickSelect(arr, q+1, rightIndex, targetIndex)
	}
	return quickSelect(arr, 1, q-1, targetIndex)
}

func randomPartition(a []int, leftIndex, rightIndex int) int {
	splitIndex := rand.Int()%(rightIndex-leftIndex+1) + 1
	a[splitIndex], a[rightIndex] = a[rightIndex], a[splitIndex]
	return partition(a, leftIndex, rightIndex)
}

func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}
