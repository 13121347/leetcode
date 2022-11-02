package Dynamic

import (
	"fmt"
	"testing"
)

func Test_lengthOfLIS2(t *testing.T) {
	testNums := []int{4, 10, 4, 3}
	test2 := lengthOfLIS(testNums)
	//test1 := lengthOfLIS(testNums)
	fmt.Println(test2)
	//fmt.Println(test1)
}
