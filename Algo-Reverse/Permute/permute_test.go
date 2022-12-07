package Permute

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	testArr := []int{1, 2, 3}
	result := permute(testArr)
	for _, v := range result {
		fmt.Println(v)
	}
}
