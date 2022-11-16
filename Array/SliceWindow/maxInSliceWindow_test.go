package SliceWindow

import (
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	testArr := []int{1, 2, 3, 1, 2, 3}
	containsNearbyDuplicate(testArr, 2)
}
