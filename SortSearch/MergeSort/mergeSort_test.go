package MergeSort

import (
	"fmt"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	toSortedArr := []int{4, 5, 3, 7, 6}
	mergeSort(toSortedArr, 0, len(toSortedArr)-1)
	fmt.Println(toSortedArr)
}
