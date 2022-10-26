package Sort

import (
	"fmt"
	"testing"
)

func TestQuickSortDown(t *testing.T) {
	testArr := []int{5, 2, 3, 8, 4, 9}
	QuickSortUp(testArr, 0, 5)
	for _, k := range testArr {
		fmt.Print(k, " ")
	}
	fmt.Println()
	fmt.Println("***************")
	QuickSortDown(testArr, 0, 5)
	for _, k := range testArr {
		fmt.Print(k, " ")
	}

}
