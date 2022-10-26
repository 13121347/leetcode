package KthLargestNumber

import (
	"fmt"
	"testing"
)

func Test_randomPartition(t *testing.T) {
	arr := []int{1, 4, 5, 2, 7, 13}
	fmt.Println(findKthLargest(arr, 1))
}
