package SumTwoNumber

import (
	"fmt"
	"testing"
)

func TestSumTwoNum1(t *testing.T) {
	nums := []int{1, 2, 3, 7}
	target := 8
	fmt.Println(twoSum(nums, target))
}
