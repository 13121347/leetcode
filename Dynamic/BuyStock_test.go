package Dynamic

import (
	"fmt"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	testArr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := trap(testArr)
	fmt.Println(result)
}
