package InsertSort

import (
	"fmt"
	"testing"
)

func Test_insertSort(t *testing.T) {
	toSortedArr := []int{4, 4, 5, 3, 7, 6}
	insertSort2(toSortedArr)
	fmt.Println(toSortedArr)
}
