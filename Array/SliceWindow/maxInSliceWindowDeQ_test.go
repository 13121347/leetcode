package SliceWindow

import (
	"fmt"
	"testing"
)

func Test_maxSlidingWindowDQ(t *testing.T) {
	testArr := []int{1, 3, 5, 6, 1, 7, 4, 2}
	result := maxSlidingWindowDQ(testArr, 3)
	fmt.Println(result)
}
