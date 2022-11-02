package SliceWindow

import (
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	testArr := []int{4, 2, 3, 1, 3, 5}
	maxSlidingWindow(testArr, 3)
}
